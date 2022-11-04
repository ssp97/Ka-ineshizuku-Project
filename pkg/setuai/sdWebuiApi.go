package setuai

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"strconv"
)

type Sdtxt2imgReq struct {
	EnableHr          bool     `json:"enable_hr"`
	DenoisingStrength int      `json:"denoising_strength"`
	FirstphaseWidth   int      `json:"firstphase_width"`
	FirstphaseHeight  int      `json:"firstphase_height"`
	Prompt            string   `json:"prompt"`
	Styles            []string `json:"styles"`
	Seed              int      `json:"seed"`
	//Subseed           int      `json:"subseed"`
	//SubseedStrength   int      `json:"subseed_strength"`
	//SeedResizeFromH   int      `json:"seed_resize_from_h"`
	//SeedResizeFromW   int      `json:"seed_resize_from_w"`
	BatchSize         int      `json:"batch_size"`
	NIter             int      `json:"n_iter"`
	Steps             int      `json:"steps"`
	CfgScale          int      `json:"cfg_scale"`
	Width             int      `json:"width"`
	Height            int      `json:"height"`
	RestoreFaces      bool     `json:"restore_faces"`
	Tiling            bool     `json:"tiling"`
	NegativePrompt    string   `json:"negative_prompt"`
	//Eta               int      `json:"eta"`
	//SChurn            int      `json:"s_churn"`
	//STmax             int      `json:"s_tmax"`
	//STmin             int      `json:"s_tmin"`
	//SNoise            int      `json:"s_noise"`
	//OverrideSettings  struct {
	//} `json:"override_settings"`
	SamplerIndex string `json:"sampler_index"`
}


type Sdtxt2imgRsp struct {
	Images     []string `json:"images"`
	Parameters struct {
		EnableHr          bool     `json:"enable_hr"`
		DenoisingStrength float32      `json:"denoising_strength"`
		FirstphaseWidth   float32      `json:"firstphase_width"`
		FirstphaseHeight  float32      `json:"firstphase_height"`
		Prompt            string   `json:"prompt"`
		Styles            []string `json:"styles"`
		Seed              int      `json:"seed"`
		Subseed           int      `json:"subseed"`
		SubseedStrength   float32      `json:"subseed_strength"`
		SeedResizeFromH   float32      `json:"seed_resize_from_h"`
		SeedResizeFromW   float32      `json:"seed_resize_from_w"`
		BatchSize         int      `json:"batch_size"`
		NIter             int      `json:"n_iter"`
		Steps             int      `json:"steps"`
		CfgScale          float32      `json:"cfg_scale"`
		Width             int      `json:"width"`
		Height            int      `json:"height"`
		RestoreFaces      bool     `json:"restore_faces"`
		Tiling            bool     `json:"tiling"`
		NegativePrompt    string   `json:"negative_prompt"`
		Eta               float32      `json:"eta"`
		SChurn            float32      `json:"s_churn"`
		STmax             float32      `json:"s_tmax"`
		STmin             float32      `json:"s_tmin"`
		SNoise            float32      `json:"s_noise"`
		OverrideSettings  struct {
		} `json:"override_settings"`
		SamplerIndex string `json:"sampler_index"`
	} `json:"parameters"`
	Info string `json:"info"`
}


type Sdtxt2imgRspInfo struct {
	Prompt                string      `json:"prompt"`
	AllPrompts            []string    `json:"all_prompts"`
	NegativePrompt        string      `json:"negative_prompt"`
	Seed                  int         `json:"seed"`
	AllSeeds              []int       `json:"all_seeds"`
	Subseed               int64       `json:"subseed"`
	AllSubseeds           []int64     `json:"all_subseeds"`
	SubseedStrength       float64     `json:"subseed_strength"`
	Width                 int         `json:"width"`
	Height                int         `json:"height"`
	SamplerIndex          int         `json:"sampler_index"`
	Sampler               string      `json:"sampler"`
	CfgScale              float64     `json:"cfg_scale"`
	Steps                 int         `json:"steps"`
	BatchSize             int         `json:"batch_size"`
	RestoreFaces          bool        `json:"restore_faces"`
	FaceRestorationModel  interface{} `json:"face_restoration_model"`
	SdModelHash           string      `json:"sd_model_hash"`
	SeedResizeFromW       int         `json:"seed_resize_from_w"`
	SeedResizeFromH       int         `json:"seed_resize_from_h"`
	DenoisingStrength     float64     `json:"denoising_strength"`
	ExtraGenerationParams struct {
	} `json:"extra_generation_params"`
	IndexOfFirstImage int      `json:"index_of_first_image"`
	Infotexts         []string `json:"infotexts"`
	Styles            []string `json:"styles"`
	JobTimestamp      string   `json:"job_timestamp"`
	ClipSkip          int      `json:"clip_skip"`
}

func SdDecode(data []byte) ([]byte, string){
	j := Sdtxt2imgRsp{}
	err := json.Unmarshal(data, &j)
	if err!= nil{
		return nil,""
	}

	b64 := j.Images[0]
	img,err := base64.StdEncoding.DecodeString(b64)
	if err!= nil{
		return nil,""
	}

	k := Sdtxt2imgRspInfo{}
	err = json.Unmarshal([]byte(j.Info), &k)
	if err!= nil{
		return img,""
	}

	return img, fmt.Sprintf("seed:%d\r\n", k.Seed)

}

func SdRequest(url string, prompt, width, height, scale, sampler, steps, seed, uc *string)([]byte,string){
	j := Sdtxt2imgReq{
		false,
		0,
		0,
		0,
		"",
		[]string{},
		-1,
		//-1,
		//0,
		//-1,
		//-1,
		1,
		1,
		20,
		11,
		512,
		768,
		false,
		false,
		"",
		//0,
		//0,
		//0,
		//0,
		//1,
		//nil,
		"Euler a",
	}

	if prompt != nil{
		j.Prompt = *prompt
	}
	if width!=nil{
		j.Width, _ = strconv.Atoi(*width)
	}
	if height!=nil{
		j.Height, _ = strconv.Atoi(*height)
	}
	if scale!=nil{
		j.CfgScale, _ = strconv.Atoi(*scale)
	}
	if sampler!=nil{
		j.SamplerIndex = *sampler
	}
	if steps!=nil{
		j.Steps, _ = strconv.Atoi(*steps)
	}
	if seed!=nil{
		j.Seed, _ = strconv.Atoi(*seed)
	} else {
		var n uint32
		binary.Read(rand.Reader, binary.LittleEndian, &n)
		j.Seed = int(n)
		print("seed = ", j.Seed, "\r\n")
	}
	if uc!=nil{
		j.NegativePrompt = *uc
	}

	jsonBytes, err := json.Marshal(j)
	if err != nil{
		return nil, error.Error(err)
	}

	opt := grequests.RequestOptions{
		Headers: map[string]string{
			"content-type":"application/json",
		},
		JSON: jsonBytes,
	}

	r,err := grequests.Post(url+"/sdapi/v1/txt2img", &opt)
	img,txt := SdDecode(r.Bytes())

	return img, txt
}