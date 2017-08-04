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
				"Continuous prevention of green; yellow and black algae.",
				"Kills most  cases of algae within 4 to 24 hours.",
				"Clarifiers create ultra clear water.",
				"Boosts chlorine effectiveness up to 600%.",
				"Improves salt cells' chlorine output and performance.",
				"Eliminates chlorine odors & skin-eye irritation.",
			},
			"summary":    "Pooltec® 3-in-1 multitask product that kills and prevents all types of algae growth, ultra water clarity, and is a strong chlorine booster for both chlorine and salt pools. When used weekly, Pooltec® synergies with all types of chlorine to boost water quality and chlorine effectiveness. Pooltec® helps increase chlorine residuals in salt pools especially during heavy usage periods. Pooltec® also keeps pool water consistently treated, eliminating the need for cleanup products such as algaecides and clarifiers. The use of Pooltec® provides an easy, year-round pool maintenance solution.",
			"howItWorks": "Pooltecs'® environmentally safe polymers uniquely synergies with all types of sanitizes to provide superior water quality, clarity, and improve water sanitation. Unique cationic polymeric compounds function as a broad spectrum, non-oxidizing algaecide-microbiocide that kill and control the growth of microorganisms by disturbing their normal metabolic process of the living cell. This process archives a super-sanitized condition with average oxidizer level.",
			"features": []string{
				"High performance algaecide; water clarifiers, and chlorine booster.",
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
				"Eliminates need for other cleanup products such as algaecides and clarifiers.",
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
				"Kills and controls all algae, green, yellow and black.",
				"Superior effectiveness and fast results within just 4 to 24 hours for most algae.",
				"Eradicates difficult Black Algae growths within 7 to 10 days.",
				"Strongest, fastest acting algaecide available.",
				"Metal-free, non-hazardous to handle, will not foam or stain.",
			},
			"summary":    "Algatec® is a highly effective, dual-action algaecide that quickly kills all types of green and yellow algae in 4 to 24 hours. The product also eradicates stubborn black algae growths with minimal brushing. Algatec's® powerful interaction with chlorine based sanitizers help to enhance algae cleanup and removal. This is due to it's alga stat inhibitor function. Algatec® also leaves pool water sparkling clear due to its strong polymeric clarifying action.",
			"howItWorks": "Algatec's® unique cationic polymeric compound function as a broad spectrum, non-oxidizing algaecide-microbiocide that kill and control the growth of microorganisms by disturbing their normal metabolic process of the living cell. Algatec® has a powerful synergy with chlorine sanitizes which further enhance algae kill and removal. The product is metal-free, bromine-free, ammonia-free will not cause foam or stain and doesn't remove chlorine from pool.",
			"features": []string{
				"Kills and removes all types of green and yellow algae in 4 to 24 hours.",
				"Eradicates black algae-fungi growths in 7 to 10 days or fewer.",
				"Super-clarifies pool water during cleanup function.",
				"Provides the fastest cleanup action available compared to other algaecide products on the market today.",
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
			"tag":  "Superior Pool Scale-Stain-Scum Preventative Surface Cleaner",
			"points": []string{
				"Aggressively prevents scale, stain and scum buildup.",
				"Helps remove existing buildup on surfaces and equipment.",
				"Salt cell descaler, protector and prolongs cell life span.",
				"Controls calcium up to 1,000 ppm and hard silicates up to 300 ppm.",
				"Strongly sequesters iron, copper, and manganese metals",
				"Removes calcium films, calcium phosphate and waterline buildup.",
				"Long lasting protection, chlorine stable.",
				"Contains no phosphates, phosphonates or acids.",
			},
			"summary":    "The ultimate pool surface protection for superior scale, stain, and scum control and removal. New Synertec® polymer technology provides unsurpassed surface protection without phosphates and acids. Beautec® also removes existing scale buildup on tile and salt cell generators without scrubbing and use of harsh acid cleaners. Metal stains due to iron, copper, and manganese are strongly sequestered to prevent surface damage. Beautec® Scale-Stain controller is the only product available today that prevents both softer calcium and harder calcium-silicate scales.",
			"howItWorks": "Beautec's® advanced Synertec formulation combines a blend of high performance polymers that utilize a complex mechanism of threshold inhibitor, crystal modification and sequestration to accomplish effective preventative and ongoing removal of mineral scales, deposits, stains, and scum on pool surfaces. Beautec's® exceptional performance is due, in part, to its extreme chlorine stability..",
			"features": []string{
				"Provides superior scale, stain, and scum control Plus prevents hard silicate scales.",
				"Salt cell descaler and cleaner.",
				"Removes waterline calcium buildup.",
				"Protects entire pool year-round.",
				"Effectively removes plaster dust during new or remodeled pool startup.",
				"Controls calcium hardness levels up to 1000 ppm.",
				"Controls silicate levels up to 300 ppm.",
			},
			"selling": []string{
				"Provides complete control of all types of metals and scale including hard silicates.",
				"Excellent for cleaning and maintaining salt chlorine cells.",
				"Small monthly dosage builds protection with each application.",
				"Advanced Synertec® formula is both acid-free and phosphate free.",
				"Great for new pool start up, quickly removes plaster dust and protects pool for months after treatment.",
			},
			"benefits": []string{
				"Only 1-2 bottles per year for most pools, after initial startup dosage.",
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
				"Little to no chlorine or salt cell usage.",
				"Pooltec will not break down scale stain inhibitors.",
				"Protection not affected by heat, sunlight or evaporation.",
				"Outlast winterizing kits, works when chlorine is gone.",
				"Strong protection lasts from fall to spring.",
			},
			"summary": "Pooltec Fall-winter Treatment provides strong, lasting winter protection that outlast winterizing kits and other winter treatments. Pooltec gives your customers the assurance that their pool will be clean at spring opening. Pooltec is the only treatment that protects covered, mesh and even uncovered pools.",
			"points": []string{
				"All Season Algaecide.",
				"Water Clarifier.",
				"Non-Chlorine Treatment.",
				"Protection Lasts from fall to Spring.",
			},
			"howItWorks": "Start with a clean, balanced pool and if needed, shock pool. Pour 3/4 to one quart of pooltec Fall-Winter Treatment per 10,000 gallons at time of pool closing, along with other closing additives.",
		}
	case "scaletec":
		return &data{
			"name": strings.ToTitle(name),
			"img":  name,
			"sds":  name,
			"tag":  "Strong pool Surface and tile descaler plus stain remover.",
			"points": []string{
				"Effectively removes calcium buildup on pool surfaces, tile and equipment.",
				"Cleans and restores pool surfaces in just two weeks.",
				"Stain remover and sequesters iron, copper, and manganese metals.",
				"Salt cell cleaner and protector, helps optimize chlorine output.",
				"Replaces hazardous acid-washing method and costly tile cleaning.",
				"Advanced non-acidic and ultra-low phosphate Synertec formulation.",
			},
			"summary":    "High performance swimming pool calcium descaler, iron stain, scum remover and preventative. First pourable product that removes hard water buildup throughout pool surfaces, tile and all equipment without scrubbing or acids. Pour directly into pool water. Highly effective in preventing iron, copper, and manganese stains. Powerful Synertec formula provides long lasting, affordable protection year-round with just a small monthly application.",
			"howItWorks": "Scaletec Plus® unique Synertec formula involve components that remove both new and old white calcium and mixed composition scales and metal stains. This process involves a combination of super-sequestration, crystal modification - absorption, and threshold inhibition. Scales and stains are both loosen from surfaces and redissolved into pool water which are filtered away or kept in suspension.",
			"features": []string{
				"Strong mineral scale, metal stain, scum remover and preventative.",
				"Descales calcium buildup throughout pool without an acid-wash and pool draining.",
				"Cleans salt cell generator to remove and prevent calcium buildup.",
			},
			"selling": []string{
				"Scaletec Plus® is the first pourable calcium descaler and remover added directly to pool water to cleanup pool surfaces, tile, filters and all other pool equipment.",
				"Excellent pretreatment to an acid wash especially for pools with very old scale.",
				"Pool surface restoration, cleanup, and descaling results in just 2 to 4 weeks.",
				"Long lasting cleanup and preventive action due to highly stable chlorine components.",
				"Affordable, do-it-yourself cleanup requires no special equipment or knowledge.",
				"Save hundreds of dollars over acid washing and tile cleaning expense.",
			},
			"benefits": []string{
				"One to two bottles will descale an average 20,000 gallons of swimming pool water.",
				"One 64 Oz. bottle provides three monthly follow up maintenance dosages.",
				"Scaletec Plus® also descales pool filter, heater, water auto-leveler and solar units.",
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
			"tag":  "An Advanced Treatment for the Prevention of Plaster Dust, Staining and Mottling.",
			"points": []string{
				"Provides cleaner startups with 90%-100% less dust and brushing.",
				"Significant reduction of mottling discoloration on all types of finishes.",
				"Typically eliminates “hot startups” that exposes quartz and aggregates.",
				"Strongly sequesters metal ions that stain surfaces and form calcium film.",
				"Brilliant exposure of pigmented and pebble type finishes.",
				"Advanced multi-sequestrant formulation, no phosphates or acids.",
				"Reduces warranty issues, lawsuits, and replastering of damaged pools.",
				"Swimmers can use pool immediately.",
			},
			"summary": "Prevents plus removes plaster dust - Eliminates tedious surface brushing - Significantly reduces surface mottling discoloration - Exposes aggregates brilliantly with less acid",
		}
	case "protec":
		return &data{
			"name": strings.ToTitle(name),
			"img":  name,
			"sds":  name,
			"tag":  "",
			"points": []string{
				"Prevents and removes scale buildup.",
				"Controls water metals and prevents surface stains.",
				"Protects fountain water pumps.",
				"Eliminates distilled water usage.",
				"Small monthly dosage.",
				"Non-acidic.",
				"Environmentally safe.",
				"Treated water safe for birds, plants, and animals including aquatic life.",
			},
			"summary": "Protecs' powerful Synertec cleaners remove and prevent both white calcium scale-deposits and stains (due to iron and copper) on surfaces of decorative; fountains, birdbaths, tabletop fountains, and water gardens.",
			"notice":  "Should calcium persist, raise water levels to allow protec to \"soak\" deposit. Double the recommended dosage for stubborn calcium.",
		}
	case "fountec":
		return &data{
			"name": strings.ToTitle(name),
			"img":  name,
			"sds":  name,
			"points": []string{
				"Powerful green algae remover.",
				"Eradicates yellow and black algae.",
				"Removes and kills clinging or stubborn type algae with minimal brushing.",
				"Water clarifiers leaves water ultra-clear.",
				"Non-foaming and non-staining.",
				"No weekly chlorine or testing required.",
				"Prevents regrowth with small weekly dose.",
				"Safe for plants, birds, and animals to drink treated water.",
				"Not safe for aquatic life.",
			"summary": "Kills all types of algae and prevents regrowth in decorative fountains, water gardens, tabletops, and birdbaths. Fountecs' cationic polymers also super-clarify water.",
			"notice":  "Should Algae persit repeat visible algae dosage everyday, to further assist in algae and biofilm removal, shock foutain with household bleach.",
		}

	default:
		return &data{}
	}
}
