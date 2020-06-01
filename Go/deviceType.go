package main

import (
	"github.com/openebs/node-disk-manager/blockdevice"
	"io/ioutil"
	"strconv"
	"strings"
)

func GetDeviceType(bd blockdevice.BlockDevice) string {
	name := strings.Split(bd.DevPath,"/")[2]
	if len(bd.Parent) > 0 {
		return "part"
	}
	if strings.HasPrefix(name, "dm-") {
		if dmuuid := readFileAsString(bd.SysPath+"/dm/uuid"); len(dmuuid) > 0 {
			dmPrefix := strings.Split(dmuuid,"-")[0]
			if len(dmPrefix) > 0  {
				if strings.EqualFold(dmPrefix[0:4],"part") {
					return dmPrefix[0:4]
				}
				return dmPrefix
			}
		}
		return "dm"
	}
	if strings.Compare(name[0:4],"loop") == 0 {
		return "loop"
	}
	if strings.Compare(name[0:2],"md") == 0 {
		mdlevel := readFileAsString(bd.SysPath+"/md/level")
		if len(mdlevel) >0 {
			return mdlevel
		}
		return "md"
	}
	if x, err := strconv.Atoi(readFileAsString(bd.SysPath+"/device/type")); err == nil {
		switch x {
		
		}
	}
	return "disk"

}

func readFileAsString(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(bytes)
}
