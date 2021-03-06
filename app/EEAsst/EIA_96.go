package EEAsst

import (
	"errors"
	"fmt"
	"regexp"
)

var eia96DictA = map[string]float64{
	"_01":100,"_02":102,"_03":105,"_04":107,"_05":110,"_06":113,"_07":115,
	"_08":118,"_09":121,"_10":124,"_11":127,"_12":130,"_13":133,"_14":137,
	"_15":140,"_16":143,"_17":147,"_18":150,"_19":154,"_20":158,"_21":162,
	"_22":165,"_23":169,"_24":174,"_25":178,"_26":182,"_27":187,"_28":191,
	"_29":196,"_30":200,"_31":205,"_32":210,"_33":215,"_34":221,"_35":226,
	"_36":232,"_37":237,"_38":243,"_39":249,"_40":255,"_41":261,"_42":267,
	"_43":274,"_44":280,"_45":287,"_46":294,"_47":301,"_48":309,"_49":316,
	"_50":324,"_51":332,"_52":340,"_53":348,"_54":357,"_55":365,"_56":374,
	"_57":383,"_58":392,"_59":402,"_60":412,"_61":422,"_62":432,"_63":442,
	"_64":453,"_65":464,"_66":475,"_67":487,"_68":499,"_69":511,"_70":523,
	"_71":536,"_72":549,"_73":562,"_74":576,"_75":590,"_76":604,"_77":619,
	"_78":634,"_79":649,"_80":665,"_81":681,"_82":698,"_83":715,"_84":732,
	"_85":750,"_86":768,"_87":787,"_88":806,"_89":825,"_90":845,"_91":866,
	"_92":887,"_93":909,"_94":931,"_95":953,"_96":976,
}
var eia96DictB = map[string]float64{
	"Z":0.001,"Y":0.01,"R":0.01,"X":0.1,"S":0.1,"A":1,"B":10,"H":10,
	"C":100,"D":1000,"E":10000,"F":100000,
}

func getEIA96Res(str string)(string, error){

	reg := regexp.MustCompile("(\\d{2}[A-Z]{1})")
	result := reg.FindAllStringSubmatch(str,-1)

	if len(result) <= 0{
		return "", errors.New("Unknown")
	}
	str = result[0][0]

	fmt.Println(str[0:2],str[2:3])

	value1, ok1 := eia96DictA["_" + str[0:2]]
	value2, ok2 := eia96DictB[str[2:3]]
	if !ok1 || !ok2{
		return "", errors.New("Unknown")
	}
	value := value1*value2
	return fmt.Sprintf(autoUnit(value)), nil
}

