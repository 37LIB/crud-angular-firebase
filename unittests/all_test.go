
package unittests

import (
	"encoding/json"
	"testing"

	"github.com/liuzl/ling"
)

func TestAll(t *testing.T) {
	cases := []string{
		`乾隆爷的乾儿子是谁？`,
		`支持中國大陸、臺灣、香港異體字和地區習慣用詞轉換，如「裏」「裡」、「鼠標」「滑鼠」。`,
		`自然語言處理是人工智能領域中的一個重要方向。`,
		//Chinese
		"目前新车的轴距还没有公布，但是长度达到了4915mm，未来定位可能是一款中级轿车，而新车重量为1890kg。作为一款新能源车型，氢燃料的使用无疑让人惊喜，而本田公布它的最大续航为750km，可以说是非常能跑。",
		//English
		"Emerging after two hours of talks with Chinese President Xi Jinping, Trump said he doesn't fault China for taking advantage of differences between the way the two countries do business.",
		//Janpanese
		"「東京動画」東京都公式動画チャンネル。都政の仕組みや街の魅力を伝える、いつでも・どこでも・誰でも楽しめるコンテンツがここに集まる！",
		//Korean
		"트럼프 대통령의 방한은 대통령 취임 후 첫 번째 방한이자 문재인 대통령과의 세 번째 정상회담이고 미국 대통령으로서는 25년 만에 국빈 방한이라는 의미가 있었다. ",
		//Russian
		"Официальный сайт Московского государственного университета имени М.В.Ломоносова (МГУ)",
		//Thai
		"ทันทุกเหตุการณ์ข่าววันนี้ข่าวล่าสุดตรวจหวยดวงข่าวบันเทิงคอลัมน์ดังเรื่องย่อ",
		//Swedish
		"Swedish photograper Per-Anders Jörgensen and Art Director Lotta Jörgensen are the duo behind one of the most interesting Food Magazines in the world.",
		//English text
		"zhanliangliu@gmail.com,      \t\t\t .@#$!@  zliu.org 123 is one two three",
		",，.。有意思quanjian１２３",
		"，。！？【】（）％＃＠＆１２３４５６７８９０“”‘’''\"\"『』「」﹃﹄〔〕—-《》：、",
		"123hj is goo. goog1e brightliang137 liang@zliu.org",
		"1238 3.1415 -1.618 6.023e23 1e-13 1,234,234",
		`自建房2樓3室2廳1衛1廚92.00㎡戶型圖，92km到北京`,
		`one-hundred and twenty-five`,
		`山西省煤炭地质１４８勘查院煤层气工程处`,
	}

	pipeline, err := ling.DefaultNLP()
	if err != nil {
		t.Error(err)
	}
	for _, c := range cases {
		d := ling.NewDocument(c)
		err = pipeline.Annotate(d)
		if err != nil {
			t.Error(err)
		}
		t.Logf("lang  :%s", d.Lang)
		t.Logf("langs :%s", d.Langs)
		t.Logf("tokens:%s", d.Tokens)
		t.Logf("lower :%s", d.XTokens(ling.Lower))
		t.Logf("norm  :%s", d.XTokens(ling.Norm))
		t.Logf("lemma :%s", d.XTokens(ling.Lemma))
		t.Logf("unidecode :%s", d.XTokens(ling.Unidecode))
		b, err := json.Marshal(d)
		if err != nil {
			t.Error(err)
		}
		t.Log(string(b))
	}
}