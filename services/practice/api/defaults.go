package practiceapi

import "blinders/packages/collecting"

var (
	DefaultLanguageUnit = map[string][]collecting.SuggestPracticeUnitResponse{
		"en": {
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
				Word:        "Pitch Deck",
				Explain:     "A presentation created by entrepreneurs that provides an overview of their business, often used to seek funding from investors.",
				ExpandWords: []string{"presentation", "slide deck", "investment pitch", "startup pitch"},
			},
			{
				Word:        "Bootstrapping",
				Explain:     "Building a company from the ground up with little or no outside capital, relying on personal savings and revenue generated by the business.",
				ExpandWords: []string{"self-funding", "DIY funding", "self-sustaining", "frugal entrepreneurship"},
			},
			{
				Word:        "MVP (Minimum Viable Product)",
				Explain:     "The simplest version of a product that can be released to market, allowing a team to collect the maximum amount of validated learning about customers with the least effort.",
				ExpandWords: []string{"prototype", "early version", "beta product", "test product"},
			},
			{
				Word:        "Disruption",
				Explain:     "The process by which a new product or service creates a significant impact on an existing market, often displacing established businesses or practices.",
				ExpandWords: []string{"innovation", "game changer", "market revolution", "industry shake-up"},
			},
			{
				Word:        "Pivot",
				Explain:     "A fundamental change in a startup's business model, product, or strategy, usually undertaken in response to market feedback or changing circumstances.",
				ExpandWords: []string{"course correction", "strategic shift", "business model adjustment", "adaptive change"},
			},
			{
				Word:        "Acquisition",
				Explain:     "The process by which one company purchases a significant portion of or all of another company's shares to gain control of its assets, technologies, or talent.",
				ExpandWords: []string{"takeover", "merger", "buyout", "consolidation"},
			},
			{
				Word:        "Incubator",
				Explain:     "A program designed to help startups grow and succeed by providing resources such as office space, mentoring, and networking opportunities.",
				ExpandWords: []string{"accelerator", "startup hub", "innovation center", "launchpad"},
			},
			{
				Word:        "Scale-Up",
				Explain:     "The phase in a startup's lifecycle when it experiences rapid growth in revenue, customer base, and workforce, often requiring significant expansion of operations and resources.",
				ExpandWords: []string{"growth stage", "expansion phase", "scaling", "ramp-up"},
			},
		},
		"vi": {
			{
				Word:        "Doanh nhân",
				Explain:     "Một cá nhân tổ chức và điều hành doanh nghiệp hoặc doanh nghiệp, đảm nhận rủi ro tài chính lớn hơn bình thường để làm điều đó.",
				ExpandWords: []string{"người sáng tạo", "người sáng lập", "nhà doanh nhân", "nhà tầm nhìn"},
			},
			{
				Word:        "Vốn rủi ro",
				Explain:     "Nguồn vốn được cung cấp bởi các nhà đầu tư cho các công ty khởi nghiệp và doanh nghiệp nhỏ được cho là có tiềm năng tăng trưởng dài hạn.",
				ExpandWords: []string{"đầu tư", "vốn", "tiền gốc"},
			},
			{
				Word:        "Bài thuyết trình",
				Explain:     "Một bài thuyết trình được tạo ra bởi các doanh nhân cung cấp một tổng quan về doanh nghiệp của họ, thường được sử dụng để tìm nguồn vốn từ các nhà đầu tư.",
				ExpandWords: []string{"bài thuyết trình", "bộ trình bày", "bài thuyết trình đầu tư", "bài thuyết trình khởi nghiệp"},
			},
			{
				Word:        "Tự lực",
				Explain:     "Xây dựng một công ty từ đầu với ít hoặc không vốn bên ngoài, dựa vào tiết kiệm cá nhân và doanh thu được tạo ra bởi doanh nghiệp.",
				ExpandWords: []string{"tự tài trợ", "tài trợ DIY", "tự duy trì", "doanh nhân tiết kiệm"},
			},
			{
				Word:        "MVP (Sản phẩm tối thiểu khả dụng)",
				Explain:     "Phiên bản đơn giản nhất của một sản phẩm có thể ra mắt thị trường, cho phép một nhóm thu thập số lượng tối đa thông tin được xác nhận về khách hàng với sự cố gắng ít nhất.",
				ExpandWords: []string{"nguyên mẫu", "phiên bản sớm", "sản phẩm beta", "sản phẩm thử nghiệm"},
			},
			{
				Word:        "Đổi mới",
				Explain:     "Quá trình mà một sản phẩm hoặc dịch vụ mới tạo ra một tác động đáng kể đến thị trường hiện tại, thường làm thay đổi các doanh nghiệp hoặc thực hành đã được thành lập.",
				ExpandWords: []string{"sáng tạo", "thay đổi trò chơi", "cách mạng thị trường", "sự rung chuyển ngành"},
			},
			{
				Word:        "Chuyển đổi",
				Explain:     "Một sự thay đổi cơ bản trong mô hình kinh doanh, sản phẩm hoặc chiến lược của một công ty khởi nghiệp, thường được thực hiện để đáp ứng phản hồi của thị trường hoặc hoàn cảnh thay đổi.",
				ExpandWords: []string{"điều chỉnh hướng đi", "chuyển đổi chiến lược", "điều chỉnh mô hình kinh doanh", "thay đổi thích ứng"},
			},
			{
				Word:        "Sự mua lại",
				Explain:     "Quá trình trong đó một công ty mua một phần đáng kể hoặc toàn bộ cổ phần của một công ty khác để giành quyền kiểm soát tài sản, công nghệ hoặc tài năng của công ty đó.",
				ExpandWords: []string{"tiếp quản", "sáp nhập", "mua lại", "hợp nhất"},
			},
			{
				Word:        "Vườn ươm",
				Explain:     "Một chương trình được thiết kế để giúp các công ty khởi nghiệp phát triển và thành công bằng cách cung cấp các nguồn lực như không gian văn phòng, cố vấn và cơ hội kết nối.",
				ExpandWords: []string{"trung tâm khởi nghiệp", "trung tâm đổi mới", "bệ phóng khởi nghiệp"},
			},
			{
				Word:        "Mở rộng quy mô",
				Explain:     "Giai đoạn trong vòng đời của một công ty khởi nghiệp khi nó có sự tăng trưởng nhanh chóng về doanh thu, cơ sở khách hàng và lực lượng lao động, thường đòi hỏi phải mở rộng đáng kể các hoạt động và nguồn lực.",
				ExpandWords: []string{"giai đoạn tăng trưởng", "giai đoạn mở rộng", "mở rộng quy mô", "tăng tốc"},
			},
		},
	}
	DefaultExplain = map[string][]collecting.ExplainEvent{
		"en": {
			{
				Request: collecting.ExplainRequest{
					Text:     "Startups are innovative new businesses that aim to disrupt traditional industries.",
					Sentence: "Startups are innovative new businesses",
				},
				Response: collecting.ExplainResponse{
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
				Request: collecting.ExplainRequest{
					Text:     "Startup ecosystems thrive on collaboration, innovation, and rapid growth.",
					Sentence: "Startup ecosystems thrive on collaboration",
				},
				Response: collecting.ExplainResponse{
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
				Request: collecting.ExplainRequest{
					Text:     "Startups often seek funding from venture capitalists to fuel their growth.",
					Sentence: "Startups often seek funding",
				},
				Response: collecting.ExplainResponse{
					Translate: "Các startup thường tìm kiếm vốn từ các nhà đầu tư mạo hiểm để thúc đẩy sự phát triển của họ.",
					GrammarAnalysis: map[string]any{
						"tense": map[string]string{
							"type":       "present",
							"identifier": "simple",
						},
						"structure": map[string]string{
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
				Request: collecting.ExplainRequest{
					Text:     "Startups often seek funding from venture capitalists to fuel their growth.",
					Sentence: "Startups often seek funding",
				},
				Response: collecting.ExplainResponse{
					Translate: "Các startup thường tìm kiếm vốn từ các nhà đầu tư mạo hiểm để thúc đẩy sự phát triển của họ.",
					GrammarAnalysis: map[string]any{
						"tense": map[string]string{
							"type":       "present",
							"identifier": "simple",
						},
						"structure": map[string]string{
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
				Request: collecting.ExplainRequest{
					Text:     "Many startups fail due to lack of market research and poor product-market fit.",
					Sentence: "Many startups fail",
				},
				Response: collecting.ExplainResponse{
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
				Request: collecting.ExplainRequest{
					Text:     "Successful startups often have a strong vision and a clear go-to-market strategy.",
					Sentence: "Successful startups often have a strong vision",
				},
				Response: collecting.ExplainResponse{
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
				Request: collecting.ExplainRequest{
					Text:     "Startup founders need to be resilient and adaptable to navigate through challenges.",
					Sentence: "Startup founders need to be resilient",
				},
				Response: collecting.ExplainResponse{
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
				Request: collecting.ExplainRequest{
					Text:     "Startup culture often values creativity, risk-taking, and continuous learning.",
					Sentence: "Startup culture often values creativity",
				},
				Response: collecting.ExplainResponse{
					Translate: "Văn hóa startup thường đánh giá cao sự sáng tạo, sẵn lòng chấp nhận rủi ro và học hỏi liên tục.",
					GrammarAnalysis: map[string]any{
						"tense": map[string]string{
							"type":       "present",
							"identifier": "simple",
						},
						"structure": map[string]string{
							"type":      "declarative",
							"structure": "Startup culture often values + noun + noun, + noun + and + adjective + noun",
							"for":       "Describing the characteristics of startup culture",
						},
					},
					ExpandWords: []string{"creativity", "risk-taking", "continuous learning"},
					KeyWords:    []string{"startup culture", "creativity", "learning"},
				},
			},
		},
		"vi": {
			{
				Request: collecting.ExplainRequest{
					Text:     "Các startup là những doanh nghiệp mới sáng tạo nhằm vào việc làm gián đoạn các ngành công nghiệp truyền thống.",
					Sentence: "Các startup là những doanh nghiệp mới sáng tạo",
				},
				Response: collecting.ExplainResponse{
					Translate: "Startups are innovative new businesses that aim to disrupt traditional industries.",
					GrammarAnalysis: map[string]interface{}{
						"tense": map[string]interface{}{
							"type":       "present",
							"identifier": "simple",
						},
						"structure": map[string]interface{}{
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
				Request: collecting.ExplainRequest{
					Text:     "Hệ sinh thái khởi nghiệp phát triển mạnh mẽ trên sự hợp tác, sáng tạo và tăng trưởng nhanh chóng.",
					Sentence: "Hệ sinh thái khởi nghiệp phát triển mạnh mẽ trên sự hợp tác",
				},
				Response: collecting.ExplainResponse{
					Translate: "Startup ecosystems thrive on collaboration, innovation, and rapid growth.",
					GrammarAnalysis: map[string]interface{}{
						"tense": map[string]interface{}{
							"type":       "present",
							"identifier": "simple",
						},
						"structure": map[string]interface{}{
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
				Request: collecting.ExplainRequest{
					Text:     "Startups often seek funding from venture capitalists to fuel their growth.",
					Sentence: "Startups often seek funding",
				},
				Response: collecting.ExplainResponse{
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
				Request: collecting.ExplainRequest{
					Text:     "Startups often seek funding from venture capitalists to fuel their growth.",
					Sentence: "Startups often seek funding",
				},
				Response: collecting.ExplainResponse{
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
				Request: collecting.ExplainRequest{
					Text:     "Nhiều startup thất bại do thiếu nghiên cứu thị trường và không phù hợp với thị trường sản phẩm.",
					Sentence: "Nhiều startup thất bại",
				},
				Response: collecting.ExplainResponse{
					Translate: "Many startups fail due to lack of market research and poor product-market fit.",
					GrammarAnalysis: map[string]interface{}{
						"tense": map[string]interface{}{
							"type":       "present",
							"identifier": "simple",
						},
						"structure": map[string]interface{}{
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
				Request: collecting.ExplainRequest{
					Text:     "Các startup thành công thường có một tầm nhìn mạnh mẽ và một chiến lược đi ra thị trường rõ ràng.",
					Sentence: "Các startup thành công thường có một tầm nhìn mạnh mẽ",
				},
				Response: collecting.ExplainResponse{
					Translate: "Successful startups often have a strong vision and a clear go-to-market strategy.",
					GrammarAnalysis: map[string]interface{}{
						"tense": map[string]interface{}{
							"type":       "present",
							"identifier": "simple",
						},
						"structure": map[string]interface{}{
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
				Request: collecting.ExplainRequest{
					Text:     "Những người sáng lập startup cần phải kiên nhẫn và linh hoạt để vượt qua những thách thức.",
					Sentence: "Những người sáng lập startup cần phải kiên nhẫn",
				},
				Response: collecting.ExplainResponse{
					Translate: "Startup founders need to be resilient and adaptable to navigate through challenges.",
					GrammarAnalysis: map[string]interface{}{
						"tense": map[string]interface{}{
							"type":       "present",
							"identifier": "simple",
						},
						"structure": map[string]interface{}{
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
				Request: collecting.ExplainRequest{
					Text:     "Startup culture often values creativity, risk-taking, and continuous learning.",
					Sentence: "Startup culture often values creativity",
				},
				Response: collecting.ExplainResponse{
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
		},
	}
)
