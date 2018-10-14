<i class='last-modified'>last modified 15:59 October 14 2018</i>
### Convert User Struct to JSON
**Kuhles Meerschweinchen:**  
![img](https://files.gamebanana.com/img/ico/sprays/5076a41eda117.gif)

```go
func (us *User) ToJSON() []byte {
	jsondata, err := json.Marshal(us)
	if err != nil {
		fmt.Println("Error when converting", us, "to JSON")
		return nil
	}
	return jsondata
}
```
`Ctrl + B`