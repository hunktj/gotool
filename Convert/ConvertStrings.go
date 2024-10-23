package gotool

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// ToString 转换成string
func ToString(value any) string {
	result := ""
	if value == nil {
		return result
	}

	v := reflect.ValueOf(value)

	switch value.(type) {
	case float32, float64:
		result = strconv.FormatFloat(v.Float(), 'f', -1, 64)
		return result
	case int, int8, int16, int32, int64:
		result = strconv.FormatInt(v.Int(), 10)
		return result
	case uint, uint8, uint16, uint32, uint64:
		result = strconv.FormatUint(v.Uint(), 10)
		return result
	case string:
		result = v.String()
		return result
	case []byte:
		result = string(v.Bytes())
		return result
	default:
		newValue, _ := json.Marshal(value)
		result = string(newValue)
		return result
	}
}

// ToInt64 将值转换为 int64，如果输入不是数字格式，则返回 0 和错误
func ToInt64(value any) (int64, error) {
	v := reflect.ValueOf(value)
	var result int64
	err := fmt.Errorf("ToInt: invalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = v.Int()
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = int64(v.Uint())
		return result, nil
	case float32, float64:
		result = int64(v.Float())
		return result, nil
	case string:
		result, err = strconv.ParseInt(v.String(), 0, 64)
		if err != nil {
			result = 0
		}
		return result, err
	default:
		return result, err
	}
}

// ToInt 将值转换为 int/int32
func ToInt(s any) int {
	val, _ := ToInt64(s)
	return int(val)
}

// ToFloat 将值转换为 float64，如果输入不是浮点数，则返回 0.0 和错误
func ToFloat(value any) (float64, error) {
	v := reflect.ValueOf(value)
	result := 0.0
	err := fmt.Errorf("ToInt: unvalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = float64(v.Int())
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = float64(v.Uint())
		return result, nil
	case float32, float64:
		result = v.Float()
		return result, nil
	case string:
		result, err = strconv.ParseFloat(v.String(), 64)
		if err != nil {
			result = 0.0
		}
		return result, err
	default:
		return result, err
	}
}

// ToJson 将值转换为有效的 json 字符串
func ToJson(value any) (string, error) {
	result, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

// ToBytes 将接口转换为字节
func ToBytes(value any) ([]byte, error) {
	v := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		number := v.Int()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case uint, uint8, uint16, uint32, uint64:
		number := v.Uint()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case float32:
		number := float32(v.Float())
		bits := math.Float32bits(number)
		bytes := make([]byte, 4)
		binary.BigEndian.PutUint32(bytes, bits)
		return bytes, nil
	case float64:
		number := v.Float()
		bits := math.Float64bits(number)
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, bits)
		return bytes, nil
	case bool:
		return strconv.AppendBool([]byte{}, v.Bool()), nil
	case string:
		return []byte(v.String()), nil
	case []byte:
		return v.Bytes(), nil
	default:
		newValue, err := json.Marshal(value)
		return newValue, err
	}
}

// ToBool 将字符串转换为布尔值
func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// Rounding 四舍五入，ROUND_HALF_UP 模式实现
// 返回将 val 根据指定精度 precision（十进制小数点后数字的数目）进行四舍五入的结果。precision 也可以是负数或零。
func Rounding(val float64, precision int) float64 {
	if precision == 0 {
		return math.Round(val)
	}
	p := math.Pow10(precision)
	if precision < 0 {
		return math.Floor(val*p+0.5) * math.Pow10(-precision)
	}

	return math.Floor(val*p+0.5) / p
}
