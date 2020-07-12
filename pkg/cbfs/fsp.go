package cbfs

import (
	"io"
	"log"
)

func init() {
	if err := RegisterFileReader(&SegReader{Type: TypeFSP, Name: "fsp", New: NewFSP}); err != nil {
		log.Fatal(err)
	}
}

//NewFSP returns a ReadWriter interface for the CBFS type TypeFSP
func NewFSP(f *File) (ReadWriter, error) {
	rec := &FSPRecord{File: *f}
	return rec, nil
}

func (r *FSPRecord) Read(in io.ReadSeeker) error {
	return nil
}

func (r *FSPRecord) String() string {
	return recString(r.File.Name, r.RecordStart, r.Type.String(), r.Size, "none")
}

func (r *FSPRecord) Write(w io.Writer) error {
	return Write(w, r.FData)
}

//Header returns a pointer to the corresponding File
func (r *FSPRecord) Header() *File {
	return &r.File
}
