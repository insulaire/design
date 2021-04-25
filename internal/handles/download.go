package handles

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	basePath := "C:\\Users\\79087\\Downloads\\"
	fi, err := os.Open(basePath + "深入理解计算机系统(原书第三版3).pdf")
	if err != nil {
		return
	}

	defer fi.Close()
	c.Writer.Header().Add("Content-type", "application/download")
	c.Writer.Header().Add("Content-Disposition", "attachment;filename="+fi.Name())
	log.Println(c.Request.Header)
	buf := make([]byte, 512)
	for {
		_, err := fi.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println(err.Error())
			}
			return
		}
		_, err = c.Writer.Write(buf)
		if err != nil {

		}
	}
	// _, err = io.Copy(c.Writer, fi)
	// if err != nil {
	// 	return
	// }
}
