package util

import (
	"golang_server/Packet"
	"math"
	"math/rand"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
)

type number interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~float64 | ~float32
}

func Min[T number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Max[T number](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func Floor[T number](x T) int {
	return int(math.Floor(float64(x)))
}

func Ceil[T number](x T) int {
	return int(math.Ceil(float64(x)))
}

func Round[T number](x T) int {
	return int(math.Round(float64(x)))
}

func Contains[T comparable](slice []T, v T) bool {
	for _, s := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func Contains_f[T any](slice []T, fn func(T) bool) bool {
	for _, t := range slice {
		if fn(t) {
			return true
		}
	}
	return false
}

func FindFunc[T any, V comparable](slice []T, v V, fn func(T, V) bool) (T, bool) {
	for i := range slice {
		if fn(slice[i], v) {
			return slice[i], true
		}
	}
	return *new(T), false
}

func Remove[T comparable](slice []T, v T) []T {
	filtered_slice := []T{}
	for _, t := range slice {
		if t != v {
			filtered_slice = append(filtered_slice, t)
		}
	}
	return filtered_slice
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	filtered_slice := []T{}
	for _, t := range slice {
		if fn(t) {
			filtered_slice = append(filtered_slice, t)
		}
	}
	return filtered_slice
}

func GetCurrentDateMSec() int64 {
	return time.Now().UnixMilli()
}

func GetRandomValue(max_value int) int {
	//max 미포함
	return rand.Intn(max_value)
}

func GetRandomValueFloat() float64 {
	return rand.Float64()
}

func GetRandomRangeValue(min_value int, max_value int) int {
	//min, max 포함
	return GetRandomValue(max_value-min_value+1) + min_value
}

func GetRandomRangeValueFloat(min_value float64, max_value float64) float64 {
	//min, max 포함
	return (GetRandomValueFloat() * (max_value - min_value)) + min_value
}

func RandomOK(rate int, args ...int) bool {
	max_rate := 100
	if len(args) != 0 {
		max_rate = args[0]
	}
	return GetRandomValue(max_rate) < rate
}

func RandomOKFloat(rate float64) bool {
	return GetRandomValueFloat() < rate
}

func GetListRandomValue[T any](list []T) T {
	return list[GetRandomValue(len(list))]
}

func MakePacketBuffer(builder *flatbuffers.Builder, packet_code Packet.PacketCode, bodyOffset flatbuffers.UOffsetT) []byte {
	Packet.PacketBaseStart(builder)
	Packet.PacketBaseAddBodyType(builder, packet_code)
	Packet.PacketBaseAddBody(builder, bodyOffset)
	baseOffset := Packet.PacketBaseEnd(builder)
	builder.Finish(baseOffset)
	return builder.FinishedBytes()
}
