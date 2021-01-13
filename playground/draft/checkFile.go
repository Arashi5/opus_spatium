package draft

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"os"
	"sort"
	"strings"
	"time"
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
	Port     string
	User     string
	Password string
	Path     string
}

func exec() {
	cfg := &configs{}
	ftp, err := initFTPConnection(cfg)
	if err != nil {
		os.Exit(1)
	}
	checkFile(ftp, cfg)
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

func checkFile(ftp *ftp.ServerConn, cfg *configs) func() Files {
	tf := Files{}
	return func() Files {
		fmt.Println(files)
		cf := Files{}
		ff := Files{}
		e, err := ftp.List(cfg.Path)
		if err != nil {
			return cf
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
			return cf
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
			ff = append(ff, cf[i])
			tf = append(tf[:b], tf[b+1:]...)
		}

		return ff
	}
}
