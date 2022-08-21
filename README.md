# Map Exercise

การแยกคำใน string ที่ง่ายที่สุดคือการอาศัยตัวช่วยจาก package strings
สิ่งนี้สามารถหาพบได้จากการ search ด้วย keyword เริ่มต้นด้วยคำว่า `golang`
เช่น `golang split string` เป็นต้น
โดย package strings มีอย่างน้อย 2 function ที่สามารถช่วยให้เรื่องนี้ง่ายขึ้นคือ

```go
func Split(s, sep string) []string
```

และ

```go
func Fields(s string) []string
```

Split คือการสั่งให้แยก string ด้วย seperator เช่น " "
ส่วน Fields คือการแยกคำโดย white space ซึ่งกรณีของเรา
space ก็คือ white space ตัวหนึ่งเช่นกัน จำสามารถทำได้สองวิธีดังนี้

```go
s := "Apple Banana Apple Banana apple"
slice1 := strings.Split(s, " ")

slice2 := strings.Fields(s)
```

slice1 และ slice2 จะได้ผลลัพธ์เหมือนกัน
