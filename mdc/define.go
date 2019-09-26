package mdc


//  1.先按照设备区分
//  2.再按照测点区分
//  3.最后是按照时间区分
type Infos struct {
  DeviceName string
  SpotName string
  Value interface{}
  UpdateTime string
}

const(
  EMPTY = ""
)