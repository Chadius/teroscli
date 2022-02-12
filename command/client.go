package command

import (
	"bytes"
	"github.com/chadius/terosgamerules"
	"github.com/chadius/terosgameserver/rpc/github.com/chadius/teros_game_server"
	"net/http"
	"reflect"
)

// Processor processes commands.
type Processor struct {
	localRulesStrategy  terosgamerules.RulesStrategy
	remoteRulesStrategy teros_game_server.TerosGameServer
}

// NewCommandProcessor returns a new Processor with the given local and remote game rulesets.
func NewCommandProcessor(remoteGameRuleset teros_game_server.TerosGameServer, localGameRuleset terosgamerules.RulesStrategy) *Processor {
	if localGameRuleset == nil || reflect.ValueOf(localGameRuleset).IsNil() {
		localGameRuleset = &terosgamerules.GameRules{}
	}

	if remoteGameRuleset == nil || reflect.ValueOf(remoteGameRuleset).IsNil() {
		remoteGameRuleset = teros_game_server.NewTerosGameServerProtobufClient("http://localhost:8080", &http.Client{})
	}

	return &Processor{
		localRulesStrategy:  localGameRuleset,
		remoteRulesStrategy: remoteGameRuleset,
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

// GetRemoteRuleset is a getter
func (p Processor) GetRemoteRuleset() teros_game_server.TerosGameServer {
	return p.remoteRulesStrategy
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
