package draft

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

var files Files

type Files []FileMeta

type FileMeta struct {
	Size uint64
	Time time.Time
	Name string
}

type configs struct {
	Host     string
	Port     uint64
	User     string
	Password string
	Path     string
}

func (repository) ftpFileChecker() {
	cfg := &configs{}
	f, err := initFTPConnection(cfg)
	if err != nil {
		os.Exit(1)
	}
	cf := checkFile(f, cfg)
	if r := cf(); !r {
		os.Exit(1)
	}
}
func initFTPConnection(cfg *configs) (*ftp.ServerConn, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	conn, err := ftp.Dial(addr)
	if err != nil {
		return nil, err
	}

	err = conn.Login(cfg.User, cfg.Password)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func checkFile(ftp *ftp.ServerConn, cfg *configs) func() bool {
	tf := Files{}
	return func() bool {
		cf := Files{}
		e, err := ftp.List(cfg.Path)
		if err != nil {
			return false
		}

		for i := range e {
			if !strings.HasSuffix(e[i].Name, ".csv") {
				continue
			}
			cf = append(cf, FileMeta{
				Name: e[i].Name,
				Time: e[i].Time,
				Size: e[i].Size,
			})
		}

		if len(cf) < 1 {
			return false
		}

		for i := range cf {
			b := sort.Search(len(tf), func(inx int) bool { return cf[i].Name <= tf[inx].Name })
			if !(b < len(tf)) || tf[b].Name != cf[i].Name {
				tf = append(tf, cf[i])
				continue
			}

			if tf[b].Time != cf[i].Time && tf[b].Size != cf[i].Size {
				tf[b].Time = cf[i].Time
				tf[b].Size = cf[i].Size
				continue
			}

			c := sort.Search(len(files), func(inx int) bool { return cf[i].Name <= files[inx].Name })
			if c < len(files) && files[c] == cf[b] {
				continue
			}

			files = append(files, cf[i])
			tf = append(tf[:b], tf[b+1:]...)
		}

		return len(files) > 1
	}
}
