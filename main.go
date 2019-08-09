//Читаем файл выгрузки адресов из ЦХДПА
//Формируем выгрузку в файл csv адресов по г.Липецку по формату
//Индекс ОПС 
//Номер доставочного 
//Улица
//Номер дома
//Номер подъезда
//Число квартир в подъезде
//
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
//	"strings"
)

var dN [][]string

func main() {

	//os.Args[1], os.Args[2]
	flile_in, flile_out := path()

	//загрузим файлы в массивы
	d1 := loadCSV(flile_in)
        //Формат текстового файла flile_in с разделителем «;»
        //•0 Регион
        //•1 Район
        //•2 Тип населенного пункта
        //•3 Наименование населенного пункта
        //•4 Тип уровня в населенном пункте
        //•5 Наименование уровня в населенном пункте
        //•6 Тип улицы
        //•7 Название улицы
        //•8 Дом (№, Литера, Дробь, Корпус, Строение)
        //•9 Подъезд
        //•10 Квартира
        //•11 Индекс
        //•12 Номер ДУ"
        //*13 ID
	f, err := os.Create(flile_out)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", flile_out, err.Error())
	}
	defer f.Close()
        cmp1 :=""
	cmp2 :=""
        cm1 :=""
	cm2 :=""
        var idx2,du2,street2,home2,pd2,flat2 string

	for _, line := range d1 {
                //flat = line[10] //номер квартиры
	        cmp1 = line[11]+line[12]+line[7]+line[8]+line[9]+line[10]
		if line[11][0:4] == "3980"  || line[11][0:4] == "3989" {
                       if cmp1 != cmp2 { // пограничный переход 1 - изменился адрес улица,номер дома, номер подъезда, и кол-во квартир
                               cm1 = line[11]+line[12]+line[7]+line[8]+line[9]
                               if cm1 != cm2 { //пограничный переход 2 - изменился адрес,улица,номер дома,подъезд
        				fmt.Fprintf(f, "%s;ДУ-%s;%s;Дом %s;Подъезд-%s;Квартир-%s\n",idx2,du2,street2,home2,pd2,flat2)

				// Запомним составной ключ адреса (улица,номер дома,подъезд) для его применения на следующем шаге цикла
        	                        cm2 = line[11]+line[12]+line[7]+line[8]+line[9]
                                }
				// Запомним все данные адреса и составной ключ для анализа и применения на следующем шаге цикла
                                cmp2 = line[11]+line[12]+line[7]+line[8]+line[9]+line[10]
                               	idx2,du2,street2,home2,pd2,flat2 =  line[11], line[12], line[7], line[8],line[9],line[10]
                                //if line[10]!="" {
                                //fmt.Println(line[10])
				//}
			}

		}


	}
}


func loadCSV(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", path, err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = ';'            //разделитель полей в файле scv
	r.Comment = '#'          // символ комментария
	r.LazyQuotes = true      // разрешить ковычки в полях
	rows, err := r.ReadAll() //прочитать весь файл в массив [][]string
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	return rows
}

func path() (string, string) {
	if len(os.Args) < 3 {
		return "adrbd.txt", "adrbd398.csv"
	}
	return os.Args[1], os.Args[2]
}
