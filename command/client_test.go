package command_test

import (
	"bytes"
	"errors"
	"github.com/chadius/terosgamerules"
	"github.com/chadius/terosgameserver/rpc/github.com/chadius/teros_game_server"
	"github.com/cserrant/terosCLI/command"
	"github.com/cserrant/terosCLI/rulesstrategyfakes"
	"github.com/cserrant/terosCLI/terosgameserverfakes"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestClientCallsServerSuite(t *testing.T) {
	suite.Run(t, new(ClientCallsServerSuite))
}

type ClientCallsServerSuite struct {
	suite.Suite
	scriptData          []byte
	squaddieData        []byte
	powerData           []byte
	expectedTextResults []byte

	fakeLocalGameRuleset *rulesstrategyfakes.FakeRulesStrategy
}

func (suite *ClientCallsServerSuite) SetupTest() {
	suite.scriptData = []byte(`script data goes here`)
	suite.squaddieData = []byte(`squaddie data goes here`)
	suite.powerData = []byte(`power data goes here`)

	suite.expectedTextResults = []byte(`results of action`)
	suite.fakeLocalGameRuleset = &rulesstrategyfakes.FakeRulesStrategy{
		ReplayBattleScriptStub: func(scriptFileHandle, squaddieFileHandle, powerFileHandle io.Reader, output io.Writer) error {
			output.Write(suite.expectedTextResults)
			return nil
		},
	}
}

func (suite *ClientCallsServerSuite) TestWhenDoNotUseServer_ThenLocalRulesetIsUsed() {
	// Setup
	commandProcessor := command.NewCommandProcessor(nil, suite.fakeLocalGameRuleset)
	var rulesetPrintout bytes.Buffer

	// Act
	commandProcessor.ApplyRulesetToData(&command.RulesetArguments{
		ScriptData:    suite.scriptData,
		SquaddieData:  suite.squaddieData,
		PowerData:     suite.powerData,
		OutputMessage: &rulesetPrintout,
	})

	// Require
	require := require.New(suite.T())
	require.Equal(1, suite.fakeLocalGameRuleset.ReplayBattleScriptCallCount(), "Client was not called")
	require.Equal(suite.expectedTextResults, rulesetPrintout.Bytes(), "output received from mock object is different")
}

func (suite *ClientCallsServerSuite) TestWhenLocalGameRulesReturnError_ThenRaiseError() {
	// Setup
	errorLocalGameRuleset := &rulesstrategyfakes.FakeRulesStrategy{
		ReplayBattleScriptStub: func(scriptFileHandle, squaddieFileHandle, powerFileHandle io.Reader, output io.Writer) error {
			return errors.New("irrelevant error")
		},
	}

	commandProcessor := command.NewCommandProcessor(nil, errorLocalGameRuleset)
	var rulesetPrintout bytes.Buffer

	// Act
	errorCaught := commandProcessor.ApplyRulesetToData(&command.RulesetArguments{
		ScriptData:    suite.scriptData,
		SquaddieData:  suite.squaddieData,
		PowerData:     suite.powerData,
		OutputMessage: &rulesetPrintout,
	})

	// Require
	require := require.New(suite.T())
	require.Equal(1, errorLocalGameRuleset.ReplayBattleScriptCallCount(), "Client was not called")
	require.Error(errorCaught, "expected an error")
	require.Equal("irrelevant error", errorCaught.Error(), "error message is different")
	require.Empty(rulesetPrintout.Bytes(), "no output should have been generated")
}

type InjectGameRulesetSuite struct {
	suite.Suite
}

func TestInjectGameRulesetSuite(t *testing.T) {
	suite.Run(t, new(InjectGameRulesetSuite))
}

func (suite *InjectGameRulesetSuite) TestWhenNoLocalGameRulesetIsInjected_ThenUsesDefaultObject() {
	// Setup
	productionClient := &terosgamerules.GameRules{}

	// Act
	commandProcessor := command.NewCommandProcessor(nil, nil)

	// Assert
	require := require.New(suite.T())
	require.Equal(
		reflect.TypeOf(commandProcessor.GetLocalRuleset()),
		reflect.TypeOf(productionClient),
	)
}

func (suite *InjectGameRulesetSuite) TestWhenLocalGameRulesetIsInjected_ThenUsesGivenObject() {
	// Setup
	injectedRuleset := &rulesstrategyfakes.FakeRulesStrategy{}

	// Act
	commandProcessor := command.NewCommandProcessor(nil, injectedRuleset)

	// Assert
	require := require.New(suite.T())
	require.Equal(
		reflect.TypeOf(commandProcessor.GetLocalRuleset()),
		reflect.TypeOf(injectedRuleset),
	)
}

func (suite *InjectGameRulesetSuite) TestWhenRemoteGameRulesetIsInjected_ThenUsesGivenObject() {
	// Setup
	injectedRuleset := &terosgameserverfakes.FakeTerosGameServer{}

	// Act
	commandProcessor := command.NewCommandProcessor(injectedRuleset, nil)

	// Assert
	require := require.New(suite.T())
	require.Equal(
		reflect.TypeOf(commandProcessor.GetRemoteRuleset()),
		reflect.TypeOf(injectedRuleset),
	)
}

func (suite *InjectGameRulesetSuite) TestWhenNoRemoteRemoteGameRulesetIsInjected_ThenUsesDefaultObject() {
	//Setup
	productionClient := teros_game_server.NewTerosGameServerProtobufClient("http://localhost:8080", &http.Client{})

	//Act
	commandProcessor := command.NewCommandProcessor(nil, nil)

	//Assert
	require := require.New(suite.T())
	require.Equal(
		reflect.TypeOf(commandProcessor.GetRemoteRuleset()),
		reflect.TypeOf(productionClient),
	)
}

type LocalPackageErrorReportingSuite struct {
	suite.Suite
	fakeLocalImageTransformerClient rulesstrategyfakes.FakeRulesStrategy
}

func TestLocalPackageErrorReportingSuite(t *testing.T) {
	suite.Run(t, new(LocalPackageErrorReportingSuite))
}

func (suite *LocalPackageErrorReportingSuite) TestWhenLocalPackageReturnsError_ThenReturnError() {
	// Setup
	suite.fakeLocalImageTransformerClient = rulesstrategyfakes.FakeRulesStrategy{
		ReplayBattleScriptStub: func(inputImageDataByteStream, formulaDataByteStream, outputSettingsDataByteStream io.Reader, output io.Writer) error {
			return errors.New("irrelevant error")
		},
	}

	var dummyOutputImageData bytes.Buffer

	// Act
	commandProcessor := command.NewCommandProcessor(nil, &suite.fakeLocalImageTransformerClient)
	err := commandProcessor.ApplyRulesetToData(&command.RulesetArguments{
		ScriptData:    nil,
		SquaddieData:  nil,
		PowerData:     nil,
		OutputMessage: &dummyOutputImageData,
	})

	// Require
	require := require.New(suite.T())
	require.Error(err, "No error was found")
	require.Equal("irrelevant error", err.Error(), "error message does not match")
}
