
type obj map[string]interface{}

func setObject(name string) *obj {
	switch name {

	case "pooltec":
		return &obj{
			"name": strings.ToTitle(name),
			"tag":  "Pooltec® 3-in-1 Pool Water Treatment",
			"points": []string{
				"Continuous prevention of green, yellow and black algae. Kills most algae in 4 to 24 hours.",
				"Strong clarifiers create ultra-clear water.",
				"Boosts chlorine’s effectiveness up to 6X (600%).",
				"Improves salt cell’s chlorine output and performance.",
				"Eliminates chlorine odors & skin-eye irritation.",
			},
			"summary":    "pooltec® Pooltec® 3-in-1 multi-task product that kills and prevents all types of algae growth, ultra clarifies water, and is a strong chlorine booster for both chlorine and salt pools. used weekly, Pooltec® synergizes with all types of chlorine to boost water quality and chlorine effectiveness. It helps increase chlorine residuals in salt pools especially during heavy use periods. Pooltec® also keeps pool water consistantly treated, eliminating cleanup products such as algaecides and clarifiers. Overall, Pooltec® provides for easier, year-round pool maintenance.",
			"howItWorks": "Pooltec's® enviromentally safe polymers uniquely synergize with all types of sanitizers to provide superior water quality, greater clarity, and improve water sanitation. Unique cationic polymeric compounds function as a broad spectrum, non-oxiding algaecide-microbicide that kill and control the growth of microorganisms by disturbing their normal metabolic process of the living cell. This process achives a super-sanitzed condition with average oxidizer level",
			"features": []string{
				"High performance algaecide, water clarifier, and chlorine booster.",
				"Kills and inhibits the regrowth of yellow and black type algae (fungi).",
				"Strong chlorine synergy reduces pool shocking and chlorine usage by 25 - 65%.",
				"Reduces salt cell generator operating time and increase life span of salt cell.",
				"Boosts chlorine effectiveness up to 600%.",
				"Reduces overall weekly treatment costs.",
			},
			"selling": []string{
				"The only 3-way pool enhancer treatment available that kills, clarifies amd boosts.",
				"Small weekly dosage, only 4 - 6 oz. per 10,000 gallons.",
				"Pays for itself in chlorine savings.",
				"Provides Superior treatment results without additional testing.",
				"Substantially improves water quality and swimmer comfort.",
				"Eliminates need for other cleanup products such as algaecides and clarifiers.",
				"Superior results compared to enzymes, phosphate removers, brimides, silver, and copper treatments.",
			},
			"benefits": []string{
				"One bottle will last the average 20,000 gallon pool twoo months.",
				"Easier overall pool care maintenance and eliminates confusing cleanup treatments.",
				"Functions extremely well with salt generator units.",
				"Easy-to-use, just pour in.",
				"No Wait, swim immediately after treatment.",
			},
			"notice": "For heavily used or commercial pools, use maximum weekly dosage for better results.",
		}
	case "algatec":
		return &obj{
			"name": strings.ToTitle(name),
			"tag":  "Fast Effective Algae Removal",
			"points": []string{
				"Kills and Controls all algae, green, yellow and black.",
				"Superior effectiveness and fast results in just 4 to 24 hours - Most Algae.",
				"Eradicates difficult Black Algae(fungi) growths 7 to 10 day.",
				"Strongest, fastest acting algaecide available.",
				"Metal-free, non-hazardous to handle, will not foam or stain.",
			},
			"summary":    "Algatec® is a highly effective, dual-action algaecide that quickly kills all types of green and yellow algae in 4 to 24 hours. The product also eradicates stubborn black algae-fungi growths with minimal brushing. Algatec's® powerful interaction with clorine based sanitizers help to enhance algae cleanup and removal. This is due to it's algastat inhibitor function. Algatec® also leaves pool water sparkling clear due to it's strong polymeric clarifying action.",
			"howItWorks": "Algatec's® unique cationic polymeric compound function as a broad spectrum, non-oxiding algaecide-microbicide that kill and control the growth of microorganisms by distubing their normal metabolic process of the living cell. Algatec® has a powerful synergy with chlorine sanitizers which further enhance algae kill and removal. the product is metal-free, bromide-free, ammonia-free will not cause foam or stain and doesn't remove chlorine from pool.",
			"features": []string{
				"kills and removes all types of green and yellow algae in 4 to 24 hours.",
				"Eradicates black algae-fungi growths in 7 to 10 days or less.",
				"Super-clarifies pool water during cleanup function.",
				"Provides the fastest cleanup action available compared to other algaecide's on the market today.",
			},
			"selling": []string{
				"Extremely effective algaecide for all pool types and enviroments.",
				"Fast acting with superior results. Strong synergy with chlorine enhances cleanup.",
				"Ultra-clarifies pool water during cleanup.",
				"Salt pool friendly,  will not interfere or deposit on cell plates.",
				"Metal-free, non-hazardous to handle, will not foam or stain.",
			},
			"benefits": []string{
				"Easy to use for algae removal and cleanup.",
				"No testing required.",
				"Will not interfere with pool water pH, chemistries or sanitizer levels.",
				"No wait time, pool is immediately ready for uses",
				"Minimal brushing required for yellow or black algae removal.",
				"Will not stain pool surface's, swimmer's hair and clothing.",
			},
			"notice": "For severe cleanup pool's, use extra chlorine to help remove dead algae.",
		}
	case "beautec":
		return &obj{
			"name": strings.ToTitle(name),
			"tag":  "Fast Effective Algae Removal",
			"points": []string{
				"Kills and Controls all algae, green, yellow and black.",
				"Superior effectiveness and fast results in just 4 to 24 hours - Most Algae.",
				"Eradicates difficult Black Algae(fungi) growths 7 to 10 day.",
				"Strongest, fastest acting algaecide available.",
				"Metal-free, non-hazardous to handle, will not foam or stain.",
			},
			"summary":    "Algatec® is a highly effective, dual-action algaecide that quickly kills all types of green and yellow algae in 4 to 24 hours. The product also eradicates stubborn black algae-fungi growths with minimal brushing. Algatec's® powerful interaction with clorine based sanitizers help to enhance algae cleanup and removal. This is due to it's algastat inhibitor function. Algatec® also leaves pool water sparkling clear due to it's strong polymeric clarifying action.",
			"howItWorks": "Algatec's® unique cationic polymeric compound function as a broad spectrum, non-oxiding algaecide-microbicide that kill and control the growth of microorganisms by distubing their normal metabolic process of the living cell. Algatec® has a powerful synergy with chlorine sanitizers which further enhance algae kill and removal. the product is metal-free, bromide-free, ammonia-free will not cause foam or stain and doesn't remove chlorine from pool.",
			"features": []string{
				"kills and removes all types of green and yellow algae in 4 to 24 hours.",
				"Eradicates black algae-fungi growths in 7 to 10 days or less.",
				"Super-clarifies pool water during cleanup function.",
				"Provides the fastest cleanup action available compared to other algaecide's on the market today.",
			},
			"selling": []string{
				"Provides complete control of all types of metal.",
				"Fast acting with superior results. Strong synergy with chlorine enhances cleanup.",
				"Ultra-clarifies pool water during cleanup.",
				"Salt pool friendly,  will not interfere or deposit on cell plates.",
				"Metal-free, non-hazardous to handle, will not foam or stain.",
			},
			"benefits": []string{
				"Only 1 to 2 bottles per year for most pools, after initial startup dosage.",
				"It's easy and affordable to use, just add monthly or twice per year.",
				"Removes existing buildup without costly tile cleaning and acid washing.",
				"Easy to use, just pour in and swim immediately",
			},
			"notice": "",
		}
	default:
		return &obj{}
	}
}