package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"golang.org/x/term"
)

func main() {
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	reader := bufio.NewReader(os.Stdin)

	const NL = 13
	const BS = 127
	const LEFT = "279168"
	const RIGHT = "279167"
	const START = 0
	const NEXT = 1
	const SIZE = 9

	var SKILLS = [...]string{"create", "become", "enter", "leave", "take", "drop", "move", "transform", "warp", "learn", "use", "script", "follow"}
	var force string
	var old string
	var primitive string
	var use bool
	var key rune
	var keys string
	var scripts []string
	var left string

	force = "ghost"
	vessels := make([]string, 1)
	pockets := make(map[string][]string)
	users := make(map[string][]string)

	vessels[START] = "ghost"

	for { //two data structures forming one thought algorithms - one code data formed thought the combination of two data codes thought a code data. tie then thought controls
		fmt.Println("PRIMITIVE : ", primitive)
		fmt.Println("VESSELS   : ", vessels)
		fmt.Println("POCKETS   : ", pockets)
		fmt.Println("USERS     : ", users)
		fmt.Println("FORCE     : ", force)
		fmt.Println("SCRIPTS   : ", scripts)
		fmt.Println("PROMPT    : ", left)
		fmt.Println(" - - - - - - - - -")

		if use != true {
			key, _, _ = reader.ReadRune()

			if key != NL && key != BS {
				keys = keys + fmt.Sprintf("%d", int(key))
			}

			if reader.Buffered() > START {
				continue
			}
		}

		if key != NL && use != true {
			if keys == LEFT {
				if len(left) == START && len(primitive) != START {
					left = primitive[:len(primitive)]
				} else if len(primitive) != START {
					left = left[:len(left)-NEXT]
				}

				keys = ""
				continue
			}

			if keys == RIGHT {
				if len(left) != len(primitive) && len(primitive) != START {
					left = primitive[:len(left)+NEXT]
				} else if len(primitive) != START {
					left = left[:START]
				}

				keys = ""
				continue
			}

			if key != BS {
				left = left + string(key)
				primitive = left + primitive[len(left)-NEXT:]

				keys = ""
			} else if len(primitive) != START && key == BS {
				if len(left) != START {
					primitive = left[:len(left)-NEXT] + primitive[len(left):]
					left = left[:len(left)-NEXT]
				} else {
					primitive = left[:len(left)] + primitive[len(left):]
					left = left[:len(left)]
				}
			}
		} else {
			left = ""
			use = false

			if len(primitive) >= len("create") && primitive[:len("create")] == "create" {
				vessel := strings.LastIndex(primitive, " ")

				if !slices.Contains(vessels, primitive[vessel+NEXT:]) {
					vessels = append(vessels, primitive[vessel+NEXT:])

					for holder, pocket := range pockets {
						if slices.Contains(pocket, force) {
							pockets[holder] = append(pockets[holder], primitive[vessel+NEXT:])
						}
					}
				}
			}

			if len(primitive) >= len("become") && primitive[:len("become")] == "become" {
				vessel := strings.LastIndex(primitive, " ")

				if slices.Contains(vessels, primitive[vessel+NEXT:]) {
					top := true

					for _, pocket := range pockets {
						prs := slices.Contains(pocket, primitive[vessel+NEXT:])
						prs2 := slices.Contains(pocket, force)

						if prs && prs2 {
							force = primitive[vessel+NEXT:]
							top = false
						} else if !prs && prs2 || prs && !prs2 {
							top = false
						}
					}

					if top == true {
						force = primitive[vessel+NEXT:]
					}
				}
			}

			if len(primitive) >= len("enter") && primitive[:len("enter")] == "enter" {
				vessel := strings.LastIndex(primitive, " ")

				if slices.Contains(vessels, primitive[vessel+NEXT:]) && primitive[vessel+NEXT:] != force {
					top := true

					for holder, pocket := range pockets {
						prs := slices.Contains(pocket, primitive[vessel+NEXT:])
						prs2 := slices.Contains(pocket, force)

						if prs && prs2 {
							pockets[primitive[vessel+NEXT:]] = append(pockets[primitive[vessel+NEXT:]], force)
							pockets[holder] = slices.Delete(pockets[holder], slices.Index(pocket, force), slices.Index(pocket, force)+NEXT)

							top = false
						} else if !prs && prs2 || prs && !prs2 {
							top = false
						}
					}

					if top == true {
						pockets[primitive[vessel+NEXT:]] = append(pockets[primitive[vessel+NEXT:]], force)
					}
				}
			}

			if primitive == "leave" {
				for holder, pocket := range pockets {
					prs := slices.Contains(pocket, force)

					if prs {
						if len(pockets[holder]) != NEXT {
							pockets[holder] = slices.Delete(pockets[holder], slices.Index(pocket, force), slices.Index(pocket, force)+NEXT)
						} else {
							delete(pockets, holder)
						}

						vessel := holder

						for holder, pocket = range pockets {
							prs := slices.Contains(pocket, vessel)

							if prs {
								pockets[holder] = append(pockets[holder], force)
							}
						}
					}
				}
			}

			if len(primitive) >= len("take") && primitive[:len("take")] == "take" {
				vessel := strings.LastIndex(primitive, " ")

				if slices.Contains(vessels, primitive[vessel+NEXT:]) && primitive[vessel+NEXT:] != force {
					top := true

					for holder, pocket := range pockets {
						prs := slices.Contains(pocket, primitive[vessel+NEXT:])
						prs2 := slices.Contains(pocket, force)

						if prs && prs2 {
							pockets[force] = append(pockets[force], primitive[vessel+NEXT:])
							pockets[holder] = slices.Delete(pockets[holder], slices.Index(pocket, primitive[vessel+NEXT:]), slices.Index(pocket, primitive[vessel+NEXT:])+NEXT)

							top = false
						} else if !prs && prs2 || prs && !prs2 {
							top = false
						}
					}

					if top == true {
						pockets[force] = append(pockets[force], primitive[vessel+NEXT:])
					}
				}
			}

			if len(primitive) >= len("drop") && primitive[:len("drop")] == "drop" {
				vessel := strings.LastIndex(primitive, " ")
				prs := slices.Contains(pockets[force], primitive[vessel+NEXT:])

				if prs {
					if len(pockets[force]) != NEXT {
						pockets[force] = slices.Delete(pockets[force], slices.Index(pockets[force], primitive[vessel+NEXT:]), slices.Index(pockets[force], primitive[vessel+NEXT:])+NEXT)
					} else {
						delete(pockets, force)
					}

					for holder, pocket := range pockets {
						if slices.Contains(pocket, force) {
							pockets[holder] = append(pockets[holder], primitive[vessel+NEXT:])
						}
					}
				}
			}

			if len(primitive) >= len("move") && primitive[:len("move")] == "move" {
				destination := strings.LastIndex(primitive, " ")
				vessel := ""
				target := false

				for _, vessel = range vessels {
					target = strings.Contains(primitive, vessel)

					if vessel != primitive[destination+NEXT:] && target == true {
						break
					}
				}

				if slices.Contains(vessels, primitive[destination+NEXT:]) && primitive[destination+NEXT:] != force && target == true && primitive[destination+NEXT:] != vessel {
					top := true

					for holder, pocket := range pockets {
						prs := slices.Contains(pocket, primitive[destination+NEXT:])
						prs2 := slices.Contains(pocket, vessel)

						if prs && prs2 {
							pockets[primitive[destination+NEXT:]] = append(pockets[primitive[destination+NEXT:]], vessel)
							pockets[holder] = slices.Delete(pockets[holder], slices.Index(pocket, vessel), slices.Index(pocket, vessel)+NEXT)

							top = false
						} else if !prs && prs2 || prs && !prs2 {
							top = false
						}
					}

					if top == true {
						pockets[primitive[destination+NEXT:]] = append(pockets[primitive[destination+NEXT:]], vessel)
					}
				}
			}

			if len(primitive) >= len("transform") && primitive[:SIZE] == "transform" {
				vessel := strings.LastIndex(primitive, " ")

				if slices.Contains(vessels, primitive[vessel+NEXT:]) && primitive[vessel+NEXT:] != force {
					top := true

					for holder, pocket := range pockets {
						prs := slices.Contains(pocket, primitive[vessel+NEXT:])
						prs2 := slices.Contains(pocket, force)

						if prs && prs2 {
							previous := force
							force = primitive[vessel+NEXT:]
							vessels = slices.Delete(vessels, slices.Index(vessels, previous), slices.Index(vessels, previous)+NEXT)

							pockets[holder] = slices.Delete(pockets[holder], slices.Index(pockets[holder], previous), slices.Index(pockets[holder], previous)+NEXT)
							for skill := range users {
								if slices.Contains(users[skill], previous) {
									if len(users[skill]) != NEXT {
										users[skill] = slices.Delete(users[skill], slices.Index(users[skill], previous), slices.Index(users[skill], previous)+NEXT)
									} else {
										delete(users, skill)
									}
								}
							}

							top = false
						} else if !prs && prs2 || prs && !prs2 {
							top = false
						}
					}

					if top == true {
						previous := force
						force = primitive[vessel+NEXT:]
						vessels = slices.Delete(vessels, slices.Index(vessels, previous), slices.Index(vessels, previous)+NEXT)

						for skill := range users {
							if slices.Contains(users[skill], previous) {
								if len(users[skill]) != NEXT {
									users[skill] = slices.Delete(users[skill], slices.Index(users[skill], previous), slices.Index(users[skill], previous)+NEXT)
								} else {
									delete(users, skill)
								}
							}
						}
					}
				} else if !slices.Contains(vessels, primitive[vessel+NEXT:]) || primitive[vessel+NEXT:] == force {
					previous := force
					force = primitive[vessel+NEXT:]
					vessels = slices.Delete(vessels, slices.Index(vessels, previous), slices.Index(vessels, previous)+NEXT)
					vessels = append(vessels, force)

					for holder, pocket := range pockets {
						if slices.Contains(pocket, previous) {
							pockets[holder] = slices.Delete(pockets[holder], slices.Index(pockets[holder], previous), slices.Index(pockets[holder], previous)+NEXT)
							pockets[holder] = append(pockets[holder], force)
						}
					}

					for holder := range pockets {
						if holder == previous {
							for range pockets[holder] {
								for pos, vessel := range vessels {
									if slices.Contains(pockets[holder], vessel) {
										vessels = slices.Delete(vessels, pos, pos+NEXT)

										for skill := range users {
											if slices.Contains(users[skill], vessel) {
												if len(users[skill]) != NEXT {
													users[skill] = slices.Delete(users[skill], slices.Index(users[skill], previous), slices.Index(users[skill], previous)+NEXT)
												} else {
													delete(users, skill)
												}
											}
										}
									}
								}
							}
						}
					}

					for skill := range users {
						if slices.Contains(users[skill], previous) {
							if len(users[skill]) != NEXT {
								users[skill] = slices.Delete(users[skill], slices.Index(users[skill], previous), slices.Index(users[skill], previous)+NEXT)
							} else {
								delete(users, skill)
							}
						}
					}

					delete(pockets, previous)
				}
			}

			if len(primitive) >= len("warp") && primitive[:len("warp")] == "warp" {
				vessel := strings.LastIndex(primitive, " ")

				if slices.Contains(vessels, primitive[vessel+NEXT:]) && primitive[vessel+NEXT:] != force && !slices.Contains(pockets[primitive[vessel+NEXT:]], force) {
					pockets[primitive[vessel+NEXT:]] = append(pockets[primitive[vessel+NEXT:]], force)

					for holder := range pockets {
						if slices.Contains(pockets[holder], force) && len(pockets[holder]) != NEXT && holder != primitive[vessel+NEXT:] {
							pockets[holder] = slices.Delete(pockets[holder], slices.Index(pockets[holder], force), slices.Index(pockets[holder], force)+NEXT)
						} else if slices.Contains(pockets[holder], force) && len(pockets[holder]) == NEXT && holder != primitive[vessel+NEXT:] {
							delete(pockets, holder)
						}
					}
				} else if primitive[vessel+NEXT:] == "library" {
					for holder := range pockets {
						if slices.Contains(pockets[holder], force) && len(pockets[holder]) != NEXT {
							pockets[holder] = slices.Delete(pockets[holder], slices.Index(pockets[holder], force), slices.Index(pockets[holder], force)+NEXT)
						} else if slices.Contains(pockets[holder], force) && len(pockets[holder]) == NEXT {
							delete(pockets, holder)
						}
					}
				}
			}

			if len(primitive) >= len("learn") && primitive[:len("learn")] == "learn" {
				for _, skill := range SKILLS {
					if strings.Contains(primitive[len("learn")+NEXT:], skill) {
						for holder := range pockets {
							if slices.Contains(pockets[holder], force) {
								for skill = range users {
									if slices.Contains(users[skill], holder) {
										users[skill] = slices.Delete(users[skill], slices.Index(users[skill], holder), slices.Index(users[skill], holder)+NEXT)

										if len(users[skill]) == START {
											delete(users, skill)
										}
									}
								}

								users[primitive[6:]] = append(users[primitive[len("learn")+NEXT:]], holder)
							}
						}
					}
				}
			}

			if len(primitive) >= len("use") && primitive[:len("use")] == "use" {
				vessel := strings.LastIndex(primitive, " ")
				top := true

				for name := range users {
					if slices.Contains(users[name], primitive[vessel+NEXT:]) && primitive[vessel+NEXT:] != force {
						for holder := range pockets {
							prs := slices.Contains(pockets[holder], primitive[vessel+NEXT:])
							prs2 := slices.Contains(pockets[holder], force)

							if prs && prs2 {
								for _, skill := range SKILLS {
									if strings.Contains(name, skill) {
										old = force
										force = primitive[vessel+NEXT:]
										fmt.Println("what the fuck?", force)
										primitive = name[strings.Index(name, skill):]
										use = true
									}
								}

								top = false
							} else if !prs && prs2 || prs && !prs2 {
								top = false
							}
						}

						if top == true {
							for _, skill := range SKILLS {
								if strings.Contains(name, skill) {
									old = force
									force = primitive[vessel+NEXT:]
									primitive = name[strings.Index(name, skill):]
									use = true
								}
							}
						}
					}
				}
			}

			if len(primitive) >= len("script") && primitive[:len("script")] == "script" {
				if len(scripts) == START {
					scripts = append(scripts, primitive[len("script")+NEXT:])
				} else {
					if strings.Contains(primitive, ":") {
						prs := false

						for pos, script := range scripts {
							if strings.Contains(primitive, ":") {
								if script[:strings.Index(script, ":")] == primitive[len("script")+NEXT:strings.Index(primitive, ":")] {
									prs = true
									primitive = "script " + script
									scripts = slices.Delete(scripts, pos, pos+NEXT)
								}
							}
						}

						if prs == false {
							scripts = append(scripts, primitive[len("script")+NEXT:])
						} else {
							continue
						}
					}
				}
			}

			if len(primitive) >= len("follow") && primitive[:len("follow")] == "follow" {
				action := primitive[len("follow")+NEXT:]

				for _, script := range scripts {
					if action == script[:strings.Index(script, ":")] {
						fmt.Println(script[strings.Index(script, ":")+NEXT:])
					}
				}
			}

			if use != true {
				primitive = ""
			}

			if use == false && old != "" {
				force = old
				old = ""
			}
		}

		if len(primitive) == 70 {
			break
		}
	}
}

//WELCOME THE BUG, : CREATE CREATE, SIMPE AND EASY, DONT FEAR IT, EMBRACE IT, YEAHHH I SEE IT, ITS FUCKING COOLLLLLLLLLLL. FOLLOW LOGIC, THE BUG IS FRIEND... a BUG?
//i SEE A FEATURE... NECESSARY FOR THE LOGIC OF THE FILE SYSTEM, HAHAHAHAHAHAHAHHAHA
