package common

import (
	"database/sql"
	"errors"

	//"fmt"
	"os"
	"path/filepath"
	"strings"

	qgen "github.com/Azareal/Gosora/query_gen"
)

var Attachments AttachmentStore

type MiniAttachment struct {
	ID         int
	SectionID  int
	OriginID   int
	UploadedBy int
	Path       string
	Extra      string

	Image bool
	Ext   string
}

type Attachment struct {
	ID           int
	SectionTable string
	SectionID    int
	OriginTable  string
	OriginID     int
	UploadedBy   int
	Path         string
	Extra        string

	Image bool
	Ext   string
}

type AttachmentStore interface {
	FGet(id int) (*Attachment, error)
	Get(id int) (*MiniAttachment, error)
	MiniGetList(originTable string, originID int) (alist []*MiniAttachment, err error)
	BulkMiniGetList(originTable string, ids []int) (amap map[int][]*MiniAttachment, err error)
	Add(sectionID int, sectionTable string, originID int, originTable string, uploadedBy int, path, extra string) (int, error)
	MoveTo(sectionID, originID int, originTable string) error
	MoveToByExtra(sectionID int, originTable, extra string) error
	Count() int
	CountIn(originTable string, oid int) int
	CountInPath(path string) int
	Delete(id int) error

	AddLinked(otable string, oid int) (err error)
	RemoveLinked(otable string, oid int) (err error)
}

type DefaultAttachmentStore struct {
	fget        *sql.Stmt
	get         *sql.Stmt
	getByObj    *sql.Stmt
	add         *sql.Stmt
	count       *sql.Stmt
	countIn     *sql.Stmt
	countInPath *sql.Stmt
	move        *sql.Stmt
	moveByExtra *sql.Stmt
	delete      *sql.Stmt

	replyUpdateAttachs *sql.Stmt
	topicUpdateAttachs *sql.Stmt
}

func NewDefaultAttachmentStore(acc *qgen.Accumulator) (*DefaultAttachmentStore, error) {
	a := "attachments"
	return &DefaultAttachmentStore{
		fget:        acc.Select(a).Columns("originTable, originID, sectionTable, sectionID, uploadedBy, path, extra").Where("attachID=?").Prepare(),
		get:         acc.Select(a).Columns("originID, sectionID, uploadedBy, path, extra").Where("attachID=?").Prepare(),
		getByObj:    acc.Select(a).Columns("attachID, sectionID, uploadedBy, path, extra").Where("originTable=? AND originID=?").Prepare(),
		add:         acc.Insert(a).Columns("sectionID, sectionTable, originID, originTable, uploadedBy, path, extra").Fields("?,?,?,?,?,?,?").Prepare(),
		count:       acc.Count(a).Prepare(),
		countIn:     acc.Count(a).Where("originTable=? and originID=?").Prepare(),
		countInPath: acc.Count(a).Where("path=?").Prepare(),
		move:        acc.Update(a).Set("sectionID=?").Where("originID=? AND originTable=?").Prepare(),
		moveByExtra: acc.Update(a).Set("sectionID=?").Where("originTable=? AND extra=?").Prepare(),
		delete:      acc.Delete(a).Where("attachID=?").Prepare(),

		// TODO: Less race-y attachment count updates
		replyUpdateAttachs: acc.Update("replies").Set("attachCount=?").Where("rid=?").Prepare(),
		topicUpdateAttachs: acc.Update("topics").Set("attachCount=?").Where("tid=?").Prepare(),
	}, acc.FirstError()
}

func (s *DefaultAttachmentStore) MiniGetList(originTable string, originID int) (alist []*MiniAttachment, err error) {
	rows, err := s.getByObj.Query(originTable, originID)
	defer rows.Close()
	for rows.Next() {
		a := &MiniAttachment{OriginID: originID}
		err := rows.Scan(&a.ID, &a.SectionID, &a.UploadedBy, &a.Path, &a.Extra)
		if err != nil {
			return nil, err
		}
		a.Ext = strings.TrimPrefix(filepath.Ext(a.Path), ".")
		if len(a.Ext) == 0 {
			return nil, errors.New("corrupt attachment path")
		}
		a.Image = ImageFileExts.Contains(a.Ext)
		alist = append(alist, a)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	if len(alist) == 0 {
		err = sql.ErrNoRows
	}
	return alist, err
}

func (s *DefaultAttachmentStore) BulkMiniGetList(originTable string, ids []int) (amap map[int][]*MiniAttachment, err error) {
	if len(ids) == 0 {
		return nil, sql.ErrNoRows
	}
	if len(ids) == 1 {
		res, err := s.MiniGetList(originTable, ids[0])
		return map[int][]*MiniAttachment{ids[0]: res}, err
	}

	amap = make(map[int][]*MiniAttachment)
	var buffer []*MiniAttachment
	var currentID int
	rows, err := qgen.NewAcc().Select("attachments").Columns("attachID,sectionID,originID,uploadedBy,path").Where("originTable=?").In("originID", ids).Orderby("originID ASC").Query(originTable)
	defer rows.Close()
	for rows.Next() {
		a := &MiniAttachment{}
		err := rows.Scan(&a.ID, &a.SectionID, &a.OriginID, &a.UploadedBy, &a.Path)
		if err != nil {
			return nil, err
		}
		a.Ext = strings.TrimPrefix(filepath.Ext(a.Path), ".")
		if len(a.Ext) == 0 {
			return nil, errors.New("corrupt attachment path")
		}
		a.Image = ImageFileExts.Contains(a.Ext)
		if currentID == 0 {
			currentID = a.OriginID
		}
		if a.OriginID != currentID {
			if len(buffer) > 0 {
				amap[currentID] = buffer
				currentID = a.OriginID
				buffer = nil
			}
		}
		buffer = append(buffer, a)
	}
	if len(buffer) > 0 {
		amap[currentID] = buffer
	}
	return amap, rows.Err()
}

func (s *DefaultAttachmentStore) FGet(id int) (*Attachment, error) {
	a := &Attachment{ID: id}
	err := s.fget.QueryRow(id).Scan(&a.OriginTable, &a.OriginID, &a.SectionTable, &a.SectionID, &a.UploadedBy, &a.Path, &a.Extra)
	if err != nil {
		return nil, err
	}
	a.Ext = strings.TrimPrefix(filepath.Ext(a.Path), ".")
	if len(a.Ext) == 0 {
		return nil, errors.New("corrupt attachment path")
	}
	a.Image = ImageFileExts.Contains(a.Ext)
	return a, nil
}

func (s *DefaultAttachmentStore) Get(id int) (*MiniAttachment, error) {
	a := &MiniAttachment{ID: id}
	err := s.get.QueryRow(id).Scan(&a.OriginID, &a.SectionID, &a.UploadedBy, &a.Path, &a.Extra)
	if err != nil {
		return nil, err
	}
	a.Ext = strings.TrimPrefix(filepath.Ext(a.Path), ".")
	if len(a.Ext) == 0 {
		return nil, errors.New("corrupt attachment path")
	}
	a.Image = ImageFileExts.Contains(a.Ext)
	return a, nil
}

func (s *DefaultAttachmentStore) Add(sectionID int, sectionTable string, originID int, originTable string, uploadedBy int, path, extra string) (int, error) {
	res, err := s.add.Exec(sectionID, sectionTable, originID, originTable, uploadedBy, path, extra)
	if err != nil {
		return 0, err
	}
	lid, err := res.LastInsertId()
	return int(lid), err
}

func (s *DefaultAttachmentStore) MoveTo(sectionID, originID int, originTable string) error {
	_, err := s.move.Exec(sectionID, originID, originTable)
	return err
}

func (s *DefaultAttachmentStore) MoveToByExtra(sectionID int, originTable, extra string) error {
	_, err := s.moveByExtra.Exec(sectionID, originTable, extra)
	return err
}

func (s *DefaultAttachmentStore) Count() (count int) {
	err := s.count.QueryRow().Scan(&count)
	if err != nil {
		LogError(err)
	}
	return count
}

func (s *DefaultAttachmentStore) CountIn(originTable string, oid int) (count int) {
	err := s.countIn.QueryRow(originTable, oid).Scan(&count)
	if err != nil {
		LogError(err)
	}
	return count
}

func (s *DefaultAttachmentStore) CountInPath(path string) (count int) {
	err := s.countInPath.QueryRow(path).Scan(&count)
	if err != nil {
		LogError(err)
	}
	return count
}

func (s *DefaultAttachmentStore) Delete(id int) error {
	_, err := s.delete.Exec(id)
	return err
}

// TODO: Split this out of this store
func (s *DefaultAttachmentStore) AddLinked(otable string, oid int) (err error) {
	switch otable {
	case "topics":
		_, err = s.topicUpdateAttachs.Exec(s.CountIn(otable, oid), oid)
		if err != nil {
			return err
		}
		err = Topics.Reload(oid)
	case "replies":
		_, err = s.replyUpdateAttachs.Exec(s.CountIn(otable, oid), oid)
		if err != nil {
			return err
		}
		err = Rstore.GetCache().Remove(oid)
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		return err
	}
	return nil
}

// TODO: Split this out of this store
func (s *DefaultAttachmentStore) RemoveLinked(otable string, oid int) (err error) {
	switch otable {
	case "topics":
		_, err = s.topicUpdateAttachs.Exec(s.CountIn(otable, oid), oid)
		if err != nil {
			return err
		}
		if tc := Topics.GetCache(); tc != nil {
			tc.Remove(oid)
		}
	case "replies":
		_, err = s.replyUpdateAttachs.Exec(s.CountIn(otable, oid), oid)
		if err != nil {
			return err
		}
		err = Rstore.GetCache().Remove(oid)
	}
	if err != nil {
		return err
	}
	return nil
}

// TODO: Add a table for the files and lock the file row when performing tasks related to the file
func DeleteAttachment(aid int) error {
	a, err := Attachments.FGet(aid)
	if err != nil {
		return err
	}
	err = deleteAttachment(a)
	if err != nil {
		return err
	}
	_ = Attachments.RemoveLinked(a.OriginTable, a.OriginID)
	return nil
}

func deleteAttachment(a *Attachment) error {
	err := Attachments.Delete(a.ID)
	if err != nil {
		return err
	}

	count := Attachments.CountInPath(a.Path)
	if count == 0 {
		err := os.Remove("./attachs/" + a.Path)
		if err != nil {
			return err
		}
	}

	return nil
}
