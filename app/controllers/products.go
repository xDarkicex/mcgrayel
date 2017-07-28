package controllers

import (
	"fmt"
	"strings"

	"github.com/xDarkicex/mcgrayel/helpers"
)

//Products this is here to pass a type to router
type Products helpers.Controller

func (products Products) Show(a helpers.RouterArgs) {
	name := a.Params[0].Value
	fmt.Println(name)
	data := setObject(name)
	fmt.Println(data)
	helpers.Render(a, "products/product", *data)
}

func (products Products) New(a helpers.RouterArgs) {

	helpers.Render(a, "products/new", map[string]interface{}{})
}

func (products Products) Edit(a helpers.RouterArgs) {

	helpers.Render(a, "products/edit", map[string]interface{}{})
}

func (products Products) Create(a helpers.RouterArgs) {
}

type data map[string]interface{}

func setObject(name string) *data {
	switch name {
	case "pooltec":
		return &data{
			"name": strings.ToTitle(name),
			"img":  name,
			"sds":  name,
			"tag":  "Pooltec® 3-in-1 Pool Water Treatment",
			"points": []string{
				"Continuous prevention of green; yellow and black algae. Kills most algae in 4 to 24 hours.",
				"Clarifies create ultra-clear water.",
				"Boosts chlorine effectiveness up to 6X (600%).",
				"Improves salt cells' chlorine output and performance.",
				"Eliminates chlorine odors & skin-eye irritation.",
			},
			"summary":    "Pooltec® 3-in-1 multitask product that kills and prevents all types of algae growth, ultra clarifies water, and is a strong chlorine booster for both chlorine and salt pools. When used weekly, Pooltec® synergies with all types of chlorine to boost water quality and chlorine effectiveness. Pooltec® helps increase chlorine residuals in salt pools especially during heavy usage periods. Pooltec® also keeps pool water consistently treated, eliminating the need for cleanup products such as algaecides and clarifies. The use of Pooltec® provides an easy, year-round pool maintenance solution.",
			"howItWorks": "Pooltecs'® environmentally safe polymers uniquely synergies with all types of sanitizes to provide superior water quality, clarity, and improve water sanitation. Unique cat ionic polymeric compounds function as a broad spectrum, non-oxiding algaecide-microbicide that kill and control the growth of microorganisms by disturbing their normal metabolic process of the living cell. This process archives a super-sanitzed condition with average oxidizer level",
			"features": []string{
				"High performance algaecide; water clarifies, and chlorine booster.",
				"Kills and inhibits the regrowth of yellow and black type algae(fungi).",
				"Strong chlorine synergy reduces pool shocking and chlorine usage by 25 - 65%.",
				"Reduces salt cell generator operating time and increase life span of salt cell.",
				"Boosts chlorine effectiveness up to 600%.",
				"Reduces overall weekly treatment costs.",
			},
			"selling": []string{
				"The only 3-way pool enhancer treatment available that kills, clarifies and boosts.",
				"Small weekly dosage, only 4 - 6 oz. per 10,000 gallons.",
				"Pays for itself in chlorine savings.",
				"Provides Superior treatment results without additional testing.",
				"Substantially improves water quality and swimmer comfort.",
				"Eliminates need for other cleanup products such as algaecides and clarifies.",
				"Superior results compared to enzymes, phosphate removers, bromides, silver, and copper treatments.",
			},
			"benefits": []string{
				"One bottle will last the average 20,000 gallon pool two months.",
				"Easier overall pool care maintenance and eliminates confusing cleanup treatments.",
				"Functions extremely well with salt generator units.",
				"Easy-to-use, just pour in.",
				"No Wait, swim immediately after treatment.",
			},
			"notice": "For heavily used or commercial pools, use maximum weekly dosage for better results.",
		}
	case "algatec":
		return &data{
			"name": strings.ToTitle(name),
			"img":  name,
			"sds":  name,
			"tag":  "Fast Effective Algae Removal",
			"points": []string{
				"Kills and Controls all algae, green, yellow and black.",
				"Superior effectiveness and fast results in just 4 to 24 hours - Most Algae.",
				"Eradicates difficult Black Algae(fungi) growths 7 to 10 day.",
				"Strongest, fastest acting algaecide available.",
				"Metal-free, non-hazardous to handle, will not foam or stain.",
			},
			"summary":    "Algatec® is a highly effective, dual-action algaecide that quickly kills all types of green and yellow algae in 4 to 24 hours. The product also eradicates stubborn black algae-fungi growths with minimal brushing. Algatecs'® powerful interaction with chlorine based sanitizes help to enhance algae cleanup and removal. This is due to it's alga stat inhibitor function. Algatec® also leaves pool water sparkling clear due to its strong polymeric clarifying action.",
			"howItWorks": "Agates'® unique cat ionic polymeric compound function as a broad spectrum, non-oxiding algaecide-microbicide that kill and control the growth of microorganisms by disturbing their normal metabolic process of the living cell. Algatec® has a powerful synergy with chlorine sanitizes which further enhance algae kill and removal. The product is metal-free, bromide-free, ammonia-free will not cause foam or stain and doesn't remove chlorine from pool.",
			"features": []string{
				"Kills and removes all types of green and yellow algae in 4 to 24 hours.",
				"Eradicates black algae-fungi growths in 7 to 10 days or fewer.",
				"Super-clarifies pool water during cleanup function.",
				"Provides the fastest cleanup action available compared to other algaecide's on the market today.",
			},
			"selling": []string{
				"Extremely effective algaecide for all pool types and environments.",
				"Fast acting with superior results. Strong synergy with chlorine enhances cleanup.",
				"Ultra-clarifies pool water during cleanup.",
				"Salt pool friendly, will not interfere or deposit on cell plates.",
				"Metal-free, non-hazardous to handle, will not foam or stain.",
			},
			"benefits": []string{
				"Easy to use for algae removal and cleanup.",
				"No testing required.",
				"Will not interfere with pool water pH, chemistries or sanitation levels.",
				"No wait time, pool is immediately ready for uses",
				"Minimal brushing required for yellow or black algae removal.",
				"Will not stain pool surface's, swimmer's hair and clothing.",
			},
			"notice": "For severe cleanup, use extra chlorine to help remove dead algae.",
		}
	case "beautec":
		return &data{
			"name": strings.ToTitle(name),
			"img":  name,
			"sds":  name,
			"tag":  "Fast Effective Algae Removal",
			"points": []string{
				"Kills and Controls all algae, green, yellow and black.",
				"Superior effectiveness and fast results in just 4 to 24 hours - Most Algae.",
				"Eradicates difficult Black Algae(fungi) growths 7 to 10 day.",
				"Strongest, fastest acting algaecide available.",
				"Metal-free, non-hazardous to handle, will not foam or stain.",
			},
			"summary":    "Algatec® is a highly effective, dual-action algaecide that quickly kills all types of green and yellow algae in 4 to 24 hours. The product also eradicates stubborn black algae-fungi growths with minimal brushing. Agates'® powerful interaction with chlorine based sanitizes help to enhance algae cleanup and removal. This is due to Agates'® alga stat inhibitor function. Algatec® also leaves pool water sparkling clear due to its strong polymeric clarifying action.",
			"howItWorks": "Agates'® unique cat ionic polymeric compound function as a broad spectrum, non-oxiding algaecide-microbicide that kill and control the growth of microorganisms by disturbing their normal metabolic process of the living cell. Algatec® has a powerful synergy with chlorine sanitizes which further enhance algae kill and removal. The product is metal-free, bromide-free, ammonia-free will not cause foam or stain and doesn't remove chlorine from pool.",
			"features": []string{
				"kills and removes all types of green and yellow algae in 4 to 24 hours.",
				"Eradicates black algae-fungi growths in 7 to 10 days or fewer.",
				"Super-clarifies pool water during cleanup function.",
				"Provides the fastest cleanup action available compared to other algaecide's on the market today.",
			},
			"selling": []string{
				"Provides complete control of all types of metal.",
				"Fast acting with superior results. Strong synergy with chlorine enhances cleanup.",
				"Ultra-clarifies pool water during cleanup.",
				"Salt pool friendly, will not interfere or deposit on cell plates.",
				"Metal-free, non-hazardous to handle, will not foam or stain.",
			},
			"benefits": []string{
				"Only 1-2 bottles per year for most pools, after initial start up dosage.",
				"It's easy and affordable to use, just add monthly or twice per year.",
				"Removes existing buildup without costly tile cleaning and acid washing.",
				"Easy to use, just pour in and swim immediately",
			},
			"notice": "",
		}
	case "pooltec-winter":
		return &data{
			"name": strings.ToTitle(name),
			"img":  name,
			"sds":  "pooltec",
			"tag":  "Fall-Winter Pool Treatment",
			"selling": []string{
				"Keeps pool water clean, clear and algae-free.",
				"Little or no chlorine of salt cell usage.",
				"Pooltec will not break down scale stain inhibitors.",
				"Protection not effected by heat, sunlight or evaporation.",
				"Outlast winterizing kits, works when chlorine is gone.",
				"Strong protection lasts from fall to spring",
			},
			"summary": "Pooltec Fall-winter Treatment provides strong, lasting winter protection that outlast winterizing kits and other winter treatments. Pooltec gives your customers the assurance that their pool will be clean at spring opening. Pooltec is the only treatment that protects covered, mesh and even uncovered pools.",
			"points": []string{
				"All Season Algaecide",
				"Water Clarifier",
				"Non-Chlorine Treatment",
				"Protection Lasts from fall to Spring",
			},
			"howItWorks": "Start with a clean, balanced pool and if needed, shock pool. Pour 3/4 to one quart of pooltec Fall-Winter Treatment per 10,000 gallons at time of pool closing, along with other closing additives.",
		}
	case "scalene":
		return &data{
			"name": strings.ToTitle(name),
			"img":  name,
			"sds":  name,
			"tag":  "Strong pool Surface and tile descaler plus stain remover.",
			"points": []string{
				"Effectively removes calcium buildup on pool surfaces, tile and equipment.",
				"Cleans and restores pool surfaces in just two weeks.",
				"Stain remover and sequesters iron, copper, and manganese metals.",
				"Salt cell cleaner, protector and helps optimize chlorine output.",
				"Replaces hazardous acid-washing method and costly tile cleaing.",
				"Advanced non-acidic and ultra-low phosphate Synertec formulation.",
			},
			"summary":    "High performance swimming pool calcium descaler, iron stain, and scum remover/preventive. First pour able product that removes hard water buildup throughout pool surfaces, tile and all equipment without scrubbing or acids. Pour directly into pool water. Highly effective in preventing iron, copper, and manganese stains. Powerful Synertec formula provides long lasting, affordable protection year-round with just a small monthly application.",
			"howItWorks": "Scalene Plus unique Synertec formula involve components that remove both new and old white calcium and mixed composition scales and metal stains. This process involves a combination of super-sequestration, crystal modification - absorption, and threshold inhibition. Scales and stains are both loosen from surfaces and redissolved into pool water hich are filtered away or kept in suspension.",
			"features": []string{
				"Strong mineral scale, metal stain, and scum remover and preventative.",
				"Descales calcium buildup throughout pool without an acid-wash and pool draining.",
				"Cleans salt cell generator to remove and prevent calcium buildup.",
			},
			"selling": []string{
				"Scalene plus is the first pour able calcium descaler-remover added directly to pool water to cleanup pool surfaces, tile, filters and all other pool equipment.",
				"Excellent pretreatment to an acid wash especially for pools with very old scale.",
				"pool surface restoration, cleanup, and descaling results in just 2 to 4 weeks.",
				"Long lasting cleanup and preventive action due to highly stable chlorine components.",
				"Affordable, do-it-yourself cleanup requires no special equipment or knowledge.",
				"Save hundreds of dollars over acid washing and tile cleaning expense.",
			},
			"benefits": []string{
				"One to two bottles will descale an average 20,000 gallons of swimming pool water.",
				"One 64 Oz. bottle provides three monthly follow up maintenance dosages.",
				"Scalene Plus also descales pool filter, heater, water auto-leveler and solar units.",
				"Non acidic, non-staining, low phosphates, and enviro-safe components.",
				"Superior winter Surface descaler-stain preventative during non-swimming season.",
			},
			"notice": "For heavily scaled pools, double the recommended cleanup dosage, and keep chlorine levels at 3.0 ppm or less. After one Month, use Beautec product to continue cleanup activity.",
		}
	case "startup-tec":
		return &data{
			"name": strings.ToTitle(name),
			"img":  name,
			"sds":  name,
		}
	case "protec":
		return &data{
			"name": strings.ToTitle(name),
			"img":  name,
			"sds":  name,
			"tag":  "",
		}
	case "fountec":
		return &data{
			"name": strings.ToTitle(name),
			"img":  name,
			"sds":  name,
		}

	default:
		return &data{}
	}
}
