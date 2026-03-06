package core

///////////////////////////////////////////////////////////////////////////
//                            Weapon Specialization Auras
///////////////////////////////////////////////////////////////////////////

func (character *Character) SwordSpecializationAura() *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:      "Sword Skill Specialization",
		BuildPhase: CharacterBuildPhaseGear,
		Duration:   NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.SwordsSkill += 3
			character.PseudoStats.TwoHandedSwordsSkill += 3
		},
	})
}

func (character *Character) AxeSpecializationAura() *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:      "Axe Skill Specialization",
		BuildPhase: CharacterBuildPhaseGear,
		Duration:   NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.AxesSkill += 3
			character.PseudoStats.TwoHandedAxesSkill += 3
		},
	})
}

func (character *Character) MaceSpecializationAura() *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:      "Mace Skill Specialization",
		BuildPhase: CharacterBuildPhaseGear,
		Duration:   NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.MacesSkill += 3
			character.PseudoStats.TwoHandedMacesSkill += 3
		},
	})
}

func (character *Character) DaggerSpecializationAura() *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:      "Dagger Skill Specialization",
		BuildPhase: CharacterBuildPhaseGear,
		Duration:   NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.DaggersSkill += 3
		},
	})
}

func (character *Character) FistWeaponSpecializationAura() *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:      "Fist Weapon Skill Specialization",
		BuildPhase: CharacterBuildPhaseGear,
		Duration:   NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.UnarmedSkill += 3
		},
	})
}

func (character *Character) PoleWeaponSpecializationAura() *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:      "Pole Weapon Skill Specialization",
		BuildPhase: CharacterBuildPhaseGear,
		Duration:   NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.StavesSkill += 3
			character.PseudoStats.PolearmsSkill += 3
		},
	})
}

func (character *Character) GunSpecializationAura() *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:      "Gun Skill Specialization",
		BuildPhase: CharacterBuildPhaseGear,
		Duration:   NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.GunsSkill += 3
		},
	})
}

func (character *Character) BowSpecializationAura() *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:      "Bow Skill Specialization",
		BuildPhase: CharacterBuildPhaseGear,
		Duration:   NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.BowsSkill += 3
		},
	})
}

func (character *Character) CrossbowSpecializationAura() *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:      "Crossbow Skill Specialization",
		BuildPhase: CharacterBuildPhaseGear,
		Duration:   NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.CrossbowsSkill += 3
		},
	})
}

func (character *Character) ThrownSpecializationAura() *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:      "Thrown Skill Specialization",
		BuildPhase: CharacterBuildPhaseGear,
		Duration:   NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.ThrownSkill += 3
		},
	})
}

func (character *Character) FeralCombatSpecializationAura() *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:      "Feral Combat Skill Specialization",
		BuildPhase: CharacterBuildPhaseGear,
		Duration:   NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.FeralCombatSkill += 3
		},
	})
}
