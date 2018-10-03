package cbfs

import (
	"io"
	"log"
)

func init() {
	if err := RegisterFileReader(&SegReader{Type: TypeBootBlock, Name: "CBFSBootBlock", New: NewBootBlock}); err != nil {
		log.Fatal(err)
	}
}

func NewBootBlock(f *File) (ReadWriter, error) {
	r := &BootBlockRecord{File: *f}
	Debug("Got header %s", r.String())
	return r, nil
}

func (r *BootBlockRecord) Read(in io.ReadSeeker) error {
	return nil
}

func (r *BootBlockRecord) String() string {
	return recString(r.File.Name, r.RecordStart, r.Type.String(), r.Size, "none")
}

func (r *BootBlockRecord) Write(w io.Writer) error {
	return Write(w, r.FData)
}

func (r *BootBlockRecord) Header() *File {
	return &r.File
}
