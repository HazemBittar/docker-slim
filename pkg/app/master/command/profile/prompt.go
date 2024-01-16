package profile

import (
	"github.com/c-bata/go-prompt"

	"github.com/slimtoolkit/slim/pkg/app/master/command"
)

var CommandSuggestion = prompt.Suggest{
	Text:        Name,
	Description: Usage,
}

var CommandFlagSuggestions = &command.FlagSuggestions{
	Names: []prompt.Suggest{
		{Text: command.FullFlagName(command.FlagTarget), Description: command.FlagTargetUsage},
		{Text: command.FullFlagName(command.FlagPull), Description: command.FlagPullUsage},
		{Text: command.FullFlagName(command.FlagShowPullLogs), Description: command.FlagShowPullLogsUsage},
		{Text: command.FullFlagName(command.FlagShowContainerLogs), Description: command.FlagShowContainerLogsUsage},
		{Text: command.FullFlagName(command.FlagEnableMondelLogs), Description: command.FlagEnableMondelLogsUsage},
		{Text: command.FullFlagName(command.FlagCRORuntime), Description: command.FlagCRORuntimeUsage},
		{Text: command.FullFlagName(command.FlagCROHostConfigFile), Description: command.FlagCROHostConfigFileUsage},
		{Text: command.FullFlagName(command.FlagCROSysctl), Description: command.FlagCROSysctlUsage},
		{Text: command.FullFlagName(command.FlagCROShmSize), Description: command.FlagCROShmSizeUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbeOff), Description: command.FlagHTTPProbeOffUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbe), Description: command.FlagHTTPProbeUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbeCmd), Description: command.FlagHTTPProbeCmdUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbeCmdFile), Description: command.FlagHTTPProbeCmdFileUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbeStartWait), Description: command.FlagHTTPProbeStartWaitUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbeRetryCount), Description: command.FlagHTTPProbeRetryCountUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbeRetryWait), Description: command.FlagHTTPProbeRetryWaitUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbePorts), Description: command.FlagHTTPProbePortsUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbeFull), Description: command.FlagHTTPProbeFullUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbeExitOnFailure), Description: command.FlagHTTPProbeExitOnFailureUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbeCrawl), Description: command.FlagHTTPProbeCrawlUsage},
		{Text: command.FullFlagName(command.FlagHTTPCrawlMaxDepth), Description: command.FlagHTTPCrawlMaxDepthUsage},
		{Text: command.FullFlagName(command.FlagHTTPCrawlMaxPageCount), Description: command.FlagHTTPCrawlMaxPageCountUsage},
		{Text: command.FullFlagName(command.FlagHTTPCrawlConcurrency), Description: command.FlagHTTPCrawlConcurrencyUsage},
		{Text: command.FullFlagName(command.FlagHTTPMaxConcurrentCrawlers), Description: command.FlagHTTPMaxConcurrentCrawlersUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbeAPISpec), Description: command.FlagHTTPProbeAPISpecUsage},
		{Text: command.FullFlagName(command.FlagHTTPProbeAPISpecFile), Description: command.FlagHTTPProbeAPISpecFileUsage},
		{Text: command.FullFlagName(command.FlagPublishPort), Description: command.FlagPublishPortUsage},
		{Text: command.FullFlagName(command.FlagPublishExposedPorts), Description: command.FlagPublishExposedPortsUsage},
		{Text: command.FullFlagName(command.FlagHostExec), Description: command.FlagHostExecUsage},
		{Text: command.FullFlagName(command.FlagHostExecFile), Description: command.FlagHostExecFileUsage},
		//{Text: command.FullFlagName(command.FlagKeepPerms), Description: command.FlagKeepPermsUsage},
		{Text: command.FullFlagName(command.FlagRunTargetAsUser), Description: command.FlagRunTargetAsUserUsage},
		{Text: command.FullFlagName(command.FlagCopyMetaArtifacts), Description: command.FlagCopyMetaArtifactsUsage},
		{Text: command.FullFlagName(command.FlagRemoveFileArtifacts), Description: command.FlagRemoveFileArtifactsUsage},
		{Text: command.FullFlagName(command.FlagUser), Description: command.FlagUserUsage},
		{Text: command.FullFlagName(command.FlagEntrypoint), Description: command.FlagEntrypointUsage},
		{Text: command.FullFlagName(command.FlagCmd), Description: command.FlagCmdUsage},
		{Text: command.FullFlagName(command.FlagWorkdir), Description: command.FlagWorkdirUsage},
		{Text: command.FullFlagName(command.FlagEnv), Description: command.FlagEnvUsage},
		{Text: command.FullFlagName(command.FlagEnvFile), Description: command.FlagEnvFileUsage},
		{Text: command.FullFlagName(command.FlagLabel), Description: command.FlagLabelUsage},
		{Text: command.FullFlagName(command.FlagVolume), Description: command.FlagVolumeUsage},
		{Text: command.FullFlagName(command.FlagLink), Description: command.FlagLinkUsage},
		{Text: command.FullFlagName(command.FlagEtcHostsMap), Description: command.FlagEtcHostsMapUsage},
		{Text: command.FullFlagName(command.FlagContainerDNS), Description: command.FlagContainerDNSUsage},
		{Text: command.FullFlagName(command.FlagContainerDNSSearch), Description: command.FlagContainerDNSSearchUsage},
		{Text: command.FullFlagName(command.FlagNetwork), Description: command.FlagNetworkUsage},
		{Text: command.FullFlagName(command.FlagHostname), Description: command.FlagHostnameUsage},
		{Text: command.FullFlagName(command.FlagExpose), Description: command.FlagExposeUsage},
		//{Text: command.FullFlagName(command.FlagExcludeMounts), Description: command.FlagExcludeMountsUsage},
		//{Text: command.FullFlagName(command.FlagExcludePattern), Description: command.FlagExcludePatternUsage},
		{Text: command.FullFlagName(command.FlagMount), Description: command.FlagMountUsage},
		{Text: command.FullFlagName(command.FlagContinueAfter), Description: command.FlagContinueAfterUsage},
		{Text: command.FullFlagName(command.FlagUseLocalMounts), Description: command.FlagUseLocalMountsUsage},
		{Text: command.FullFlagName(command.FlagUseSensorVolume), Description: command.FlagUseSensorVolumeUsage},
		{Text: command.FullFlagName(command.FlagSensorIPCMode), Description: command.FlagSensorIPCModeUsage},
		{Text: command.FullFlagName(command.FlagSensorIPCEndpoint), Description: command.FlagSensorIPCEndpointUsage},
	},
	Values: map[string]command.CompleteValue{
		command.FullFlagName(command.FlagPull):                   command.CompleteBool,
		command.FullFlagName(command.FlagShowPullLogs):           command.CompleteBool,
		command.FullFlagName(command.FlagTarget):                 command.CompleteImage,
		command.FullFlagName(command.FlagShowContainerLogs):      command.CompleteBool,
		command.FullFlagName(command.FlagEnableMondelLogs):       command.CompleteBool,
		command.FullFlagName(command.FlagPublishExposedPorts):    command.CompleteBool,
		command.FullFlagName(command.FlagHTTPProbeOff):           command.CompleteBool,
		command.FullFlagName(command.FlagHTTPProbe):              command.CompleteTBool,
		command.FullFlagName(command.FlagHTTPProbeCmdFile):       command.CompleteFile,
		command.FullFlagName(command.FlagHTTPProbeFull):          command.CompleteBool,
		command.FullFlagName(command.FlagHTTPProbeExitOnFailure): command.CompleteTBool,
		command.FullFlagName(command.FlagHTTPProbeCrawl):         command.CompleteTBool,
		command.FullFlagName(command.FlagHTTPProbeAPISpecFile):   command.CompleteFile,
		command.FullFlagName(command.FlagHostExecFile):           command.CompleteFile,
		//command.FullFlagName(command.FlagKeepPerms):              command.CompleteTBool,
		command.FullFlagName(command.FlagRunTargetAsUser):     command.CompleteTBool,
		command.FullFlagName(command.FlagRemoveFileArtifacts): command.CompleteBool,
		command.FullFlagName(command.FlagNetwork):             command.CompleteNetwork,
		//command.FullFlagName(command.FlagExcludeMounts):       command.CompleteTBool,
		//command.FullFlagName(command.FlagPathPermsFile):          command.CompleteFile,
		//command.FullFlagName(command.FlagIncludePathFile):        command.CompleteFile,
		//command.FullFlagName(command.FlagIncludeShell):           command.CompleteBool,
		command.FullFlagName(command.FlagContinueAfter): command.CompleteContinueAfter,
		//command.FullFlagName(command.FlagConsoleOutput):   command.CompleteConsoleOutput,
		command.FullFlagName(command.FlagUseLocalMounts):  command.CompleteBool,
		command.FullFlagName(command.FlagUseSensorVolume): command.CompleteVolume,
		//command.FullFlagName(command.FlagKeepTmpArtifacts):       command.CompleteBool,
		command.FullFlagName(command.FlagCROHostConfigFile): command.CompleteFile,
		command.FullFlagName(command.FlagSensorIPCMode):     command.CompleteIPCMode,
	},
}
