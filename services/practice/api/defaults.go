package practiceapi

import (
	"blinders/packages/db/collectingdb"
)

var (
	DefaultSimplePracticeUnits = []SimplePracticeUnit{
		{
			Word:        "Entrepreneur",
			Explain:     "An individual who organizes and operates a business or businesses, taking on greater than normal financial risks in order to do so.",
			ExpandWords: []string{"innovator", "founder", "businessperson", "visionary"},
		},
		{
			Word:        "Venture Capital",
			Explain:     "Funding provided by investors to startup companies and small businesses that are perceived to have long-term growth potential.",
			ExpandWords: []string{"investment", "funding", "capital", "seed money"},
		},
		{
			Word:    "Pitch Deck",
			Explain: "A presentation created by entrepreneurs that provides an overview of their business, often used to seek funding from investors.",
			ExpandWords: []string{
				"presentation",
				"slide deck",
				"investment pitch",
				"startup pitch",
			},
		},
		{
			Word:    "Bootstrapping",
			Explain: "Building a company from the ground up with little or no outside capital, relying on personal savings and revenue generated by the business.",
			ExpandWords: []string{
				"self-funding",
				"DIY funding",
				"self-sustaining",
				"frugal entrepreneurship",
			},
		},
		{
			Word:        "MVP (Minimum Viable Product)",
			Explain:     "The simplest version of a product that can be released to market, allowing a team to collect the maximum amount of validated learning about customers with the least effort.",
			ExpandWords: []string{"prototype", "early version", "beta product", "test product"},
		},
		{
			Word:    "Disruption",
			Explain: "The process by which a new product or service creates a significant impact on an existing market, often displacing established businesses or practices.",
			ExpandWords: []string{
				"innovation",
				"game changer",
				"market revolution",
				"industry shake-up",
			},
		},
		{
			Word:    "Pivot",
			Explain: "A fundamental change in a startup's business model, product, or strategy, usually undertaken in response to market feedback or changing circumstances.",
			ExpandWords: []string{
				"course correction",
				"strategic shift",
				"business model adjustment",
				"adaptive change",
			},
		},
		{
			Word:        "Acquisition",
			Explain:     "The process by which one company purchases a significant portion of or all of another company's shares to gain control of its assets, technologies, or talent.",
			ExpandWords: []string{"takeover", "merger", "buyout", "consolidation"},
		},
		{
			Word:    "Incubator",
			Explain: "A program designed to help startups grow and succeed by providing resources such as office space, mentoring, and networking opportunities.",
			ExpandWords: []string{
				"accelerator",
				"startup hub",
				"innovation center",
				"launchpad",
			},
		},
		{
			Word:        "Scale-Up",
			Explain:     "The phase in a startup's lifecycle when it experiences rapid growth in revenue, customer base, and workforce, often requiring significant expansion of operations and resources.",
			ExpandWords: []string{"growth stage", "expansion phase", "scaling", "ramp-up"},
		},
	}

	ExplainLogSamples = []collectingdb.ExplainLog{
		{
			Request: collectingdb.ExplainRequest{
				Text:     "Startups are innovative new businesses that aim to disrupt traditional industries.",
				Sentence: "Startups are innovative new businesses",
			},
			Response: collectingdb.ExplainResponse{
				Translate: "Các startup là những doanh nghiệp mới sáng tạo nhằm vào việc làm gián đoạn các ngành công nghiệp truyền thống.",
				GrammarAnalysis: map[string]any{
					"tense": map[string]any{
						"type":       "present",
						"identifier": "simple",
					},
					"structure": map[string]any{
						"type":      "declarative",
						"structure": "Startups are + adjective + noun + that + verb",
						"for":       "Describing the nature of startups",
					},
				},
				ExpandWords: []string{"innovative", "disrupt", "industries"},
				KeyWords:    []string{"startups", "businesses", "innovative"},
			},
		},
		{
			Request: collectingdb.ExplainRequest{
				Text:     "Startup ecosystems thrive on collaboration, innovation, and rapid growth.",
				Sentence: "Startup ecosystems thrive on collaboration",
			},
			Response: collectingdb.ExplainResponse{
				Translate: "Hệ sinh thái khởi nghiệp phát triển mạnh mẽ trên sự hợp tác, sáng tạo và tăng trưởng nhanh chóng.",
				GrammarAnalysis: map[string]any{
					"tense": map[string]string{
						"type":       "present",
						"identifier": "simple",
					},
					"structure": map[string]string{
						"type":      "declarative",
						"structure": "Startup ecosystems thrive + on + nouns",
						"for":       "Describing the nature of startup ecosystems",
					},
				},
				ExpandWords: []string{"collaboration", "innovation", "growth"},
				KeyWords:    []string{"startup ecosystems", "thrive", "collaboration"},
			},
		},
		{
			Request: collectingdb.ExplainRequest{
				Text:     "Many startups fail due to lack of market research and poor product-market fit.",
				Sentence: "Many startups fail",
			},
			Response: collectingdb.ExplainResponse{
				Translate: "Nhiều startup thất bại do thiếu nghiên cứu thị trường và không phù hợp với thị trường sản phẩm.",
				GrammarAnalysis: map[string]any{
					"tense": map[string]string{
						"type":       "present",
						"identifier": "simple",
					},
					"structure": map[string]string{
						"type":      "declarative",
						"structure": "Many startups fail + due to + noun + and + adjective + noun",
						"for":       "Describing the common reasons for startup failure",
					},
				},
				ExpandWords: []string{"market research", "product-market fit", "failure"},
				KeyWords:    []string{"startups", "fail", "market research"},
			},
		},
		{
			Request: collectingdb.ExplainRequest{
				Text:     "Successful startups often have a strong vision and a clear go-to-market strategy.",
				Sentence: "Successful startups often have a strong vision",
			},
			Response: collectingdb.ExplainResponse{
				Translate: "Các startup thành công thường có một tầm nhìn mạnh mẽ và một chiến lược đi ra thị trường rõ ràng.",
				GrammarAnalysis: map[string]any{
					"tense": map[string]string{
						"type":       "present",
						"identifier": "simple",
					},
					"structure": map[string]string{
						"type":      "declarative",
						"structure": "Successful startups often have + adjective + noun + and + adjective + noun",
						"for":       "Describing common characteristics of successful startups",
					},
				},
				ExpandWords: []string{"vision", "go-to-market strategy", "successful"},
				KeyWords:    []string{"successful startups", "vision", "strategy"},
			},
		},
		{
			Request: collectingdb.ExplainRequest{
				Text:     "Startup founders need to be resilient and adaptable to navigate through challenges.",
				Sentence: "Startup founders need to be resilient",
			},
			Response: collectingdb.ExplainResponse{
				Translate: "Những người sáng lập startup cần phải kiên nhẫn và linh hoạt để vượt qua những thách thức.",
				GrammarAnalysis: map[string]any{
					"tense": map[string]string{
						"type":       "present",
						"identifier": "simple",
					},
					"structure": map[string]string{
						"type":      "declarative",
						"structure": "Startup founders need to be + adjective + and + adjective + to + verb + through + noun",
						"for":       "Describing the qualities required for startup founders",
					},
				},
				ExpandWords: []string{"resilient", "adaptable", "challenges"},
				KeyWords:    []string{"startup founders", "resilient", "adaptable"},
			},
		},
		{
			Request: collectingdb.ExplainRequest{
				Text:     "Startups often seek funding from venture capitalists to fuel their growth.",
				Sentence: "Startups often seek funding",
			},
			Response: collectingdb.ExplainResponse{
				Translate: "Các startup thường tìm kiếm vốn từ các nhà đầu tư mạo hiểm để thúc đẩy sự phát triển của họ.",
				GrammarAnalysis: map[string]interface{}{
					"tense": map[string]interface{}{
						"type":       "present",
						"identifier": "simple",
					},
					"structure": map[string]interface{}{
						"type":      "declarative",
						"structure": "Startups often seek + noun + from + noun",
						"for":       "Describing the common practice of startups seeking funding",
					},
				},
				ExpandWords: []string{"funding", "venture capitalists", "growth"},
				KeyWords:    []string{"startups", "funding", "venture capitalists"},
			},
		},
		{
			Request: collectingdb.ExplainRequest{
				Text:     "Startup culture often values creativity, risk-taking, and continuous learning.",
				Sentence: "Startup culture often values creativity",
			},
			Response: collectingdb.ExplainResponse{
				Translate: "Văn hóa startup thường đánh giá cao sự sáng tạo, sẵn lòng chấp nhận rủi ro và học hỏi liên tục.",
				GrammarAnalysis: map[string]interface{}{
					"tense": map[string]interface{}{
						"type":       "present",
						"identifier": "simple",
					},
					"structure": map[string]interface{}{
						"type":      "declarative",
						"structure": "Startup culture often values + noun + noun, + noun + and + adjective + noun",
						"for":       "Describing the characteristics of startup culture",
					},
				},
				ExpandWords: []string{"creativity", "risk-taking", "continuous learning"},
				KeyWords:    []string{"startup culture", "creativity", "learning"},
			},
		},
	}
)
