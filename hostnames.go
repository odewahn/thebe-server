package main

import (
	"fmt"
	"github.com/docker/docker/pkg/random"
)

var (
	left = [...]string{
		"admiring",
		"adoring",
		"agitated",
		"angry",
		"backstabbing",
		"berserk",
		"boring",
		"clever",
		"cocky",
		"compassionate",
		"condescending",
		"cranky",
		"desperate",
		"determined",
		"distracted",
		"dreamy",
		"drunk",
		"ecstatic",
		"elated",
		"elegant",
		"evil",
		"fervent",
		"focused",
		"furious",
		"gloomy",
		"goofy",
		"grave",
		"happy",
		"high",
		"hopeful",
		"hungry",
		"insane",
		"jolly",
		"jovial",
		"kickass",
		"lonely",
		"loving",
		"mad",
		"modest",
		"naughty",
		"nostalgic",
		"pensive",
		"prickly",
		"reverent",
		"romantic",
		"sad",
		"serene",
		"sharp",
		"sick",
		"silly",
		"sleepy",
		"stoic",
		"stupefied",
		"suspicious",
		"tender",
		"thirsty",
		"trusting",
	}

	// Docker, starting from 0.7.x, generates names from notable scientists and hackers.
	// Please, for any amazing man that you add to the list, consider adding an equally amazing woman to it, and vice versa.
	right = [...]string{
"rabbit", "canada_geese", "great_plover", "hammerhead_shark", "garganey_duck", "sun_star_radiata", "porpoises", "peach_faced_lovebird", "spotted_cuscus__common", "helmet_cockatoo", "southern_flying_squirrel", "ringneck_parakeet", "horseshoe_bat", "european_eel", "bat", "goliath_beetle", "tamandua", "crab_eating_mongoose", "tamandua_collared_anteater", "willow_ptarmigan", "yak", "ophiops_lizards", "himalayan_monal_pheasant", "crab_eating_opossum", "rosy_feather_star", "thick_billed_coot", "american_oystercatcher", "siren_salamander", "domestic_cat_and_kitten", "bluefin_tuna", "forked_mouse_lemur", "ivory_billed_woodpecker", "white_spoonbill", "african_spoonbill", "marsh_harrier", "red_footed_falcon", "sage_sphinx_moth_with_curly_proboscis", "black_leopard", "zebra", "silver_marmoset", "ring_tailed_lemur", "tamarin_monkey", "peruvian_spider_monkey", "woolly_monkey", "hermit_crab", "lemur", "otter_civet", "eastern_quoll_daysure", "marine_iguana", "cassowary", "thornback_cowfish", "horned_lizard", "lion", "shetland_pony", "moose", "arabian_horse", "great_white_pelican", "appaloosa_horse", "rosy_starling", "falcon", "tufted_deer", "golden_lion_tamarin", "colobus_monkey_aka_mantled_guereza", "boston_terrier", "red_gurnard_fish", "raccoons", "toy_rabbit__a_mechanical_toy", "serpent_falcon", "sea_shell", "marmoset", "rosy_feather_star", "starfish", "black_cockatoo", "elk", "sea_otter", "peccary", "short_legged_goose", "barred_bandicoot", "wood_turtle", "malay_fox_bat", "sheldrake_duck", "coyote", "squirrel_monkey", "tadpoles", "hummingbird", "ring_tailed_cat", "slender_horned_gazelle", "elephant_seal", "common_rat__aka_brown_rat", "alpaca__huacaya", "coelacanth_fish", "sawfish_shark__maybe_longtooth", "finhorse__horse", "bush_wren", "giant_green_sea_anemone", "maleo_bird", "mastiff_dog", "pacuma_toadfish", "buzzard", "magpie", "bluebird", "kangaroo_rat", "bison_with_calf", "pygmy_piculet_bird", "european_roller", "indian_civet", "honey_buzzard", "butterfly_peacock", "beech_marten", "vampire_bat", "carp", "bonapartes_gull", "garden_spider", "common_hill_myna_bird", "silver_moony_fish", "musk_deer", "forest_dormouse", "south_american_iguana", "pionus_parrot", "barbel_flying_fish", "ibex_goats", "carolina_parakeets", "dachshunds", "roseate_cockatoo", "black_headed_gull", "greylag_goose", "garter_snake", "capybara", "star_nosed_mole", "demoiselle_crane", "collie_dog", "flying_lemur", "rufous_necked_weaver", "salmon", "paradise_flycatchers", "dove", "mouse", "kiwi", "black_jaguar", "burro", "common_pheasant", "blowfish", "african_tent_tortoise", "large_indian_civet", "aardwolf", "painted_snipe", "bush_antelope", "european_golden_plover", "sand_martin_bird", "pampas_cat", "ring_dove", "whitebar_surgeonfish", "brazilian_guinea_pigs", "moray_eel", "minnows", "turquoise_parrot", "raven", "creeper", "civet", "scottish_deerhound", "violet_turaco", "pomeranian_dog", "southern_reedbuck", "frilled_lizard", "wild_ass", "brush_wallaby", "european_edible_frog", "rabbit", "grizzly_bear", "ruffed_grouse", "sheep", "roadrunner", "barking_deer", "kite", "bull", "scallop_shell", "white_cheeked_turaco", "springboks", "canadian_lynx", "tenrec", "large_claw_crab", "night_monkey", "falcon", "cockatoo", "porcelain_crab", "warty_newts", "hummingbird", "ibis", "toucan", "mandarin_duck", "caribou_reindeer", "lovebird", "crested_pigeon", "carpet_chameleon", "king_bird_of_paradise", "cock_of_the_rock", "margay_cat", "gelada_baboon", "seriema_bird", "royal_king_flycatcher", "asp", "snapper", "c_elegans_worm", "parrot", "eastern_kingbird", "kestrel", "ruff_bird", "kingfisher", "newfoundland", "butterfly", "turtle", "crested_pardalotte", "chilean_plant_cutter", "hummingbird", "gull", "locust", "opossum", "flying_locust", "crickets", "raven", "wood_duck", "armadillo", "monkey", "scops_owl", "cobra_snake", "springer_spaniel", "butterfly_blenny_fish", "flamingo", "kite", "shrimp", "eagle", "grouse", "auk", "lorikeet", "mahi_mahi", "carp", "gecko", "baboon", "cod", "wood_quails", "wallaby", "sturgeon", "purple_martin", "roseate_spoonbill", "moth_sphinx", "common_goat", "brush_tailed_rat_kangaroo_with_baby_in_its_pouch", "coral_snake", "cat", "marmot", "haddock", "lioness", "beaver", "monitor_lizard", "alpine_marmot", "wombat", "deer", "flying_fish", "armadillo", "cuttlefish", "octopus", "antelope", "gannet", "wild_boar", "gannet", "baboon", "lemur", "hooded_seal", "monkey", "phoenix_bird", "gopher", "chamois", "meadow_lark", "slow_loris", "bat", "ferret", "spiny_dogfish", "hoopoe", "cuttlefish", "mastiff", "pigeons", "gecko", "jellyfish", "english_pointer", "deer", "peach_faced_lovebird", "pipefish", "eider_duck", "pronghorn", "rhesus_monkeys", "green_lizard", "wildebeest", "loggerhead_turtle", "pig_and_piglets", "lioness_with_baby_cubs", "kea", "mousebird", "hawksbill_turtle", "pintail_duck", "brittle_starfish", "sea_urchin", "jellyfish", "stickleback_fish", "impala", "burro", "ribbon_fish", "kangaroo_nibbling_a_tree_branch", "skua_seabird", "gecko", "mouse", "beagle", "capuchin_monkey", "spotted_goby", "african_sacred_ibis", "chanting_falcon", "blackthroat_warbler", "golden_pheasant", "flying_gurnard_fish", "belgian_tervuren_sheepdog", "european_octopus", "asian_civet", "frog", "sea_sturgeon", "boll_weevil", "skunk", "manatee", "african_elephant", "orange_winged_amazon_parrot", "fossa", "lion", "hedgehog", "clydesdale", "japanese_waxwing", "hawk", "owl", "sparrow", "iguana", "greyhound", "ring_tailed_cat", "koala_bear", "alpine_chamois", "parrot", "kestrel_common", "zanick", "cape_petrel", "goatsucker__nightjar", "pheasant", "ground_squirrel", "squirrel_with_nut", "horned_owl", "shark", "polar_bear", "eskimo_dogs", "belgian_shepherd", "partridges", "glossy_starlings", "wild_canary", "long_tailed_tit", "brook_trout", "antelope", "giant_red_flying_squirrel", "gemmous_dragonet", "marmoset", "fruit_bat", "common_starling", "lime_tree_sphinx_moth", "american_marsh_hawk", "mongoose", "marbled_cat", "shrew_tenrec", "harp_seal", "grasshopper", "weasel", "roadrunner", "lynx", "hummingbird", "snail", "softshell_turtle", "little_bittern", "mollusk", "aardvark", "tragopan", "growling_lions", "osprey", "ermines", "john_dory_fish", "lesser_spotted_eagle", "tiger", "red_firefish", "domestic_chicken", "bandicoot", "flying_squirrel", "glassfish", "sand_dollar", "american_crocodile", "american_bison", "florida_panther", "alligator", "white_rabbit", "tiger", "poodles", "octopus", "passenger_pigeon", "river_otter", "pigfooted_bandicoot", "serval", "saiga", "javan_tiger", "hummingbird_moths", "red_squirrel", "tiger_lying", "gray_wolf", "desert_fox", "european_bear", "longhorn_beetle", "spider_monkey", "jellyfish", "cormorant", "european_ibex", "black_giant_squirrel", "little_pied_cormorant", "howler_monkey", "hippopotamus", "egret", "eastern_gray_squirrel", "partridge", "rhinoceros", "long_eared_owl", "barbary_ape", "golden_eagle", "ornate_chorus_frogs", "lar_gibbon", "ermine", "african_civet", "porcupine_fish", "yellow_baboon", "angora_goat", "gorgeted_bird_of_paradise", "bohemian_waxwings", "european_night_heron", "garden_dormouse", "tree_swift", "channel_billed_cuckoo", "alactaga", "praire_dog", "bank_vole", "wildcat", "moa", "jellyfish_stephalia", "california_quail", "mink", "llama", "black_lion_marmoset", "opossum", "little_owl", "stingray", "monkfish", "goldfish", "wandering_chaetodon", "bloodhounds", "irish_setter", "pearl_spotted_barbet", "green_broadbill", "wildebeest", "lapwing_bird", "tigress_with__tiger_cubs", "baby_black_rhinocerous", "cuckoo_pheasant", "musky_rat_kangaroo", "african_marsh_harrier", "english_leceister_sheep", "euploea_midamus_butterfly", "hamster_rat", "peacock_moth", "zebra_butterfly", "alpaca", "wall_gecko", "kookaburra_bird", "flying_eagle", "alaskan_plaice", "sugar_gliders", "wood_rat", "roe_deer", "tarpan_horses", "tubularia_coronata", "giraffe_baby", "giraffe", "ostriches", "catshark", "shoe_bill_umbrette", "andean_marsupial_tree_frog", "ribbed_newt", "silver_bass", "king_charles_spaniel", "hawksbill_turtle", "horned_owl", "kitten", "gorillas", "malamute_dog", "siberian_tiger_cub", "puma", "eared_seal", "tarsie", "snakeneck_turtle", "german_grey_heath_sheep", "bearded_bellbird", "white_pelican", "mosquito", "chick_and_egg", "stork", "crowned_pigeons", "lobster", "hummingbird", "macaw_parrot", "horned_screamer_bird", "penguin_rockhopper", "green_tailed_jacamar", "spotted_dalmatian", "bull_with_horns", "crested_oriole", "chimpanzee", "striped_hyena", "foxhound", "siberian_tiger", "leopard", "cape_lion", "german_shepherd", "siberian_tiger", "snow_leopard", "couching_tiger", "lion", "axolotl", "griffin_vulture", "emu_wren", "cardinal", "pine_marten", "blue_roller_bird", "greek_tortoise", "great_eared_goatsucker", "coot_birds", "superb_lyre_bird", "trojan_horse", "bushbuck", "english_setter", "donkey", "three_toed_woodpecker", "tree_porcupines", "loris", "orangutan", "turtle", "seahorse", "wolf", "lantern_fly", "bullfrog", "emu", "owl", "egyptian_goose", "sollarium_shell", "giant_anteater", "purple_capped_lory", "water_buffalo", "asian_elephant", "african_elephant", "mimic_butterfly", "house_martin", "english_sparrow", "tree_sparrow", "water_buffalo", "starfish", "brant_goose", "canada_goose", "groundhog", "song_thrush", "caspian_tern", "brittle_star", "bumpy_toad", "cuckoo_roller", "dwarf_mongoose", "leopard_ground_squirrel", "mongoose_lemur", "lizard", "american_robin", "pied_kingfisher", "woodpecker", "malayan_badger", "cockatiel", "right_whale", "shrimp", "tree_frog", "steer", "porcupine_fish", "european_merlin", "sea_lion_and_seals", "basilisk_lizard", "great_dane", "jerboas", "sulawesi_golden_owl", "trunk_fish", "shrew", "short_eared_owl_with_sleepy_eyes", "wallachian_sheep", "guinea_pigs", "mythical_griffin", "double_crested_cormorant", "corsac_fox", "palm_civet", "dodo_bird", "umbrellabird", "opah_fish", "crested_agouti", "basket_star", "yellowjacket", "garden_spider", "praying_mantis", "long_centipede", "tsetse_fly", "ladybug", "rosalie_beetle", "water_spider", "dragonfly", "cicadas", "winged_ant", "silkworm_moth", "wasp_nest", "honey_bees", "hornets", "red_ants", "methona_butterflies", "ants", "moth", "stag_beetle", "female_garden_spider", "dung_beetle", "sap_beetle", "splendor_beetle", "tarantula", "scorpion", "scarab_beetle", "moving_leaf_insect", "boll_weevil", "rat_kangaroo", "puff_adder", "rifle_bird", "puma", "herring_gull", "helmet_shrike_bird", "tailorbird", "green_woodpecker", "golden_oriole", "snow_buntings", "luna_moth", "flying_squirrel", "scrawled_butterflyfish", "rabbit", "african_wild_cat", "scallop_shell", "lesser_anteater", "green_monkeys", "barbary_sheep", "staghound", "bighorn_sheep", "greater_honeyguide", "antlion", "sea_otter", "proboscis_monkey", "red_colobus_monkey", "grey_fox", "arabian_camel", "badger", "mythical_supercilious_owl", "galapagos_land_iguana", "cuckoo_on_branch", "booted_racket_tail_hummingbird", "alpine_accentor_bird", "scarlet_minivet", "house_mouse_and_domestic_cat", "fox_terrier_dog", "blue_jay_in_flight", "narwhal", "ruby_throat_warblers", "elephant_shrew", "kangaroo", "turtle", "jersey_cow", "eastern_chipmunk", "short_tailed_pangolin", "hedgehog", "herring", "mammoth", "secretary_bird", "snipe_fish", "european_goatsucker_bird", "sea_bass", "spotted_quail", "hermit_crab", "french_angelfish", "giant_petrel_bird", "smoothhound_shark", "boatbill_heron", "darter", "bushmaster_snake", "atlantic_wreckfish", "pine_grosbeak", "guitarfish", "murex_shell", "hairtail_silvery_fish", "arctic_tern", "king_penguins", "bullhead_fish", "eastern_kingsnake", "wood_tick", "seychelles_blue_pigeon", "common_bittern", "shoveler_duck", "bullfinch", "red_fox", "two_brown_bears", "krait_snake", "common_hare", "comoro_cuckoo_roller", "four_lined_snake", "waterbuck_antelope", "tibetan_brown_bear", "giant_petrel", "stormy_petrel", "guinea_fowl", "eagle", "kingbird", "bare_throated_bellbird", "percheron_draft_horse", "argali_sheep", "frigatebird", "grey_parrot", "camel", "spotted_cuckoo", "berkshire_pig", "parrotfish", "malayan_tapir", "diana_monkey", "arctic_cod_fish", "pitta", "sixshafted_bird_of_paradise", "jerboa", "cheetah", "red_panda", "catfish", "crested_grebe", "bufflehead_duck", "flying_fish", "angelfish", "collared_titi", "flying_dragon", "sea_sponge", "jellyfish", "asiatic_wild_dog", "calling_hare", "babirusa_wild_pig", "black_swan", "toucan", "kudu_antelope", "silkworm_on_a_leaf", "daurian_pika", "elephant_shrews", "sand_grouse", "spring_hare", "golden_tailed_treeshrew", "hispaniolan_solenodon", "boa_constrictor", "python", "caiman_crocodile", "harvest_mice", "tawny_owl", "nilotic_monitor_lizard", "harpy_eagle", "conch", "african_wild_dog", "wild_goat", "nile_valley_sunbird", "sparrowhawk", "rhinoceros_beetle", "hooded_crow", "camel", "domestic_pigeon", "greyhound", "scaly_lizard", "opossum_mouse", "chaffinch_bird", "musk_shrew", "newts", "bumblebee", "blood_pheasant", "angler_fish", "black_retriever_dog", "widow_bird", "forktailed_nightjar", "pelican", "crested_shriketit", "electric_catfish", "austalian_bee_eater_bird", "sloth", "brush_tailed_possum", "hatteria_great_fringed_lizard", "tegu_lizard", "bezoar_goat", "green_crab", "striped_jackal", "german_shepherd", "sumatran_tiger", "black_crows", "merlin", "pampas_cat", "echidna", "trigger_fish", "stickleback_gar", "bulldog", "snakebird", "goosander_duck", "long_tailed_chinchilla", "butterfly", "crested_porcupine", "spotted_hyena", "slender_lorises", "jumping_jerboa", "western_quoll_daysure", "flying_fox_fruit_bat", "leaf_nosed_bat", "european_bat", "falcon", "liger", "jacana_bird", "openbill_stork", "yellowhammers", "black_fox_squirrel", "rooster_and_hens", "whip_poor_will", "foxhound", "clouded_leopard", "rose_breasted_grosbeak", "boa_constrictor", "spotted_salamander", "king_vulture", "papuan_hornbill", "great_squirrel", "agamidae_lizard", "chameleon", "salamander", "giant_squid", "land_snail", "thornback_crab", "eastern_chipmunks", "domestic_pig", "brushtail_possum", "maned_sheep", "argus_pheasant", "hyacinth_macaws", "swordfish", "crane", "sacred_ibis", "land_crab", "great_black_cockatoo", "edible_nest_swiftlets", "cashmier_goat", "doctor_fish", "shrew", "dolphin", "fruit_bat", "peacock", "pine_moth", "striped_red_mullet", "black_adler_eagle", "carolina_parrot", "archerfish", "wolverine", "toad", "snow_leopard", "wallcreepers", "tree_swift", "snake_knot", "white_breasted_nuthatch", "sand_lizard", "sea_bass", "motmot_bird", "ameiva_lizard", "addax", "gorilla_babies", "domestic_short_hair_cat", "piping_crow", "walrus", "bear_paw_clam", "banded_broadbill", "tree_shrew", "glossy_starling", "norway_rat", "italian_greyhound", "gavial_crocodile", "fly", "galapagos_tortoise", "lilac_breasted_roller", "spotted_nathura", "amazon_parrot", "bear", "oystercatcher", "dormouse_with_nut", "burbot", "great_gray_shrike", "white_stork", "leonberger_dog", "southern_ground_hornbill", "hyrax", "pika", "great_white_heron", "russian_greyhound", "rock_dove", "newfoundland_dog", "foram", "crawfish", "miniature_pinscher", "long_eared_bat", "hawksbill_sea_turtle", "puffin", "condor", "pumpkinseed_sunfish", "moor_frogs", "northern_hawk_owl", "leatherback_sea_turtle", "jellyfish", "banded_linsang", "brown_long_eared_bat", "rock_thrush", "white_pelican", "horse_mackerel", "duckbill_platypus", "least_weasel", "larks", "common_snipe", "sugar_squirrel", "hummingbird", "lioness", "nutcracker_bird", "whale_shark", "spiny_lobster", "golden_jackal", "spider_crab", "bronzewing_birds", "barred_spinefoot", "jellyfish_chrysaora", "starfish", "crab_spider", "common_nighthawk", "pencilled_lark", "thick_knee_bird", "arched_duck", "lioness", "portuguese_man_o_war", "crested_eagle", "turkey", "persian_greyhound", "eland_antelope", "common_toad", "gliding_frog", "puffer_fish", "chubby_frog", "spotted_hyena", "zebrs", "beaver", "vulture", "short_toed_eagle", "moose", "horseshoe_crab", "albatross", "box_turtle", "_banded_armadillo", "pied_avocet", "mandrill_baboon", "dorcas_gazelle", "chacma_baboon", "crested_caracara_bird", "ruddy_shelduck", "horned_frog", "common_frog", "african_bullfrog", "european_common_frog", "surinam_toad", "common_garter_snake", "electric_fan", "climbing_frog", "parakeet", "opossum", "cockatiel", "white_tailed_eagle", "bluejay", "asse_caama_fox", "giraffe", "antelope", "kudu_antelope", "scorpion_lizard", "blue_swimming_crab", "peacock", "pheasant", "kanchil_mouse_deer", "stork", "redmullett_fish", "rooster", "turtle_dove", "european_wildcat",

	}
)

// GetRandomName generates a random name from the list of adjectives and surnames in this package
// formatted as "adjective_surname". For example 'focused_turing'. If retry is non-zero, a random
// integer between 0 and 10 will be added to the end of the name, e.g `focused_turing3`
func GetRandomName(retry int) string {
	rnd := random.Rand
begin:
	name := fmt.Sprintf("%s_%s", left[rnd.Intn(len(left))], right[rnd.Intn(len(right))])
	if name == "boring_wozniak" /* Steve Wozniak is not boring */ {
		goto begin
	}

	if retry > 0 {
		name = fmt.Sprintf("%s%d", name, rnd.Intn(10))
	}
	return name
}


func generateHostName() (string) {
  return GetRandomName(0)
}
