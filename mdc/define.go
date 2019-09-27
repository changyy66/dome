package mdc

//  1.先按照设备区分
//  2.再按照测点区分
//  3.最后是按照时间区分



type SpotInfo struct {
  DeviceName string `json:"device_name"`
  SpotName string `json:"spot_name"`
  Value interface{} `json:"value"`
  UpdateTime string `json:"update_time"`
}



const(
  EMPTY = ""
  GOTIME = "2006-01-02 15:04:05"
  MODELTABLE = "templates"
)