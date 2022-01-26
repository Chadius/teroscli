package command

import (
	"bytes"
	"github.com/chadius/terosgamerules"
	"reflect"
)

// Processor processes commands.
type Processor struct {
	localRulesStrategy terosgamerules.RulesStrategy
}

// NewCommandProcessor returns a new Processor with the given local and remote game rulesets.
func NewCommandProcessor(remoteGameRuleset, localGameRuleset terosgamerules.RulesStrategy) *Processor {
	if localGameRuleset == nil || reflect.ValueOf(localGameRuleset).IsNil() {
		localGameRuleset = &terosgamerules.GameRules{}
	}

	return &Processor{
		localRulesStrategy: localGameRuleset,
	}
}

// ApplyRulesetToData uses the given arguments and applies it to the ruleset.
func (p *Processor) ApplyRulesetToData(args *RulesetArguments) error {
	return p.useLocalPackageToApplyRuleset(args)
}

func (p Processor) useLocalPackageToApplyRuleset(args *RulesetArguments) error {
	scriptDataReader := bytes.NewBuffer(args.ScriptData)
	squaddieDataReader := bytes.NewBuffer(args.SquaddieData)
	powerDataReader := bytes.NewBuffer(args.PowerData)

	return p.localRulesStrategy.ReplayBattleScript(
		scriptDataReader,
		squaddieDataReader,
		powerDataReader,
		args.OutputMessage,
	)
}

// GetLocalRuleset is a getter.
func (p *Processor) GetLocalRuleset() terosgamerules.RulesStrategy {
	return p.localRulesStrategy
}

// RulesetArguments holds the arguments scanned from the command line.
type RulesetArguments struct {
	ScriptData    []byte
	SquaddieData  []byte
	PowerData     []byte
	OutputMessage *bytes.Buffer
}
