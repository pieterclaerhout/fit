// Generated by reader_util_test.go.
// Only edit to bootstrap new entries or change existing entries.

package fit_test

var decodeTestFiles = [...]struct {
	folder      string
	name        string
	wantErr     bool
	fingerprint uint64
	compress    bool
	dopts       testingDecodeOpts
}{
	{
		"me",
		"activity-small-fenix2-run.fit",
		false,
		11891133592397655863,
		true,
		tdoAllWithNullLogger,
	},
	{
		"fitsdk",
		"Activity.fit",
		false,
		2025243915659230630,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"MonitoringFile.fit",
		false,
		13937971957850042454,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"Settings.fit",
		false,
		2003530012377202593,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WeightScaleMultiUser.fit",
		false,
		3477797669220222488,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WorkoutCustomTargetValues.fit",
		false,
		7545884481242290315,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WorkoutIndividualSteps.fit",
		false,
		2983131181881988632,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WorkoutRepeatGreaterThanStep.fit",
		false,
		1617723397621021464,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WorkoutRepeatSteps.fit",
		false,
		7028162149605874664,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WeightScaleSingleUser.fit",
		false,
		1069422916205585336,
		true,
		tdoNone,
	},
	{
		"fitsdk",
		"WeightScaleSingleUser.fit",
		false,
		1069422916205585336,
		true,
		tdoNone,
	},
	{
		"python-fitparse",
		"garmin-edge-500-activitiy.fit",
		false,
		9061108094696170205,
		true,
		tdoNone,
	},
	{
		"python-fitparse",
		"sample-activity-indoor-trainer.fit",
		false,
		1167463203483035313,
		true,
		tdoNone,
	},
	{
		"python-fitparse",
		"compressed-speed-distance.fit",
		false,
		0,
		false,
		tdoNone,
	},
	{
		"python-fitparse",
		"antfs-dump.63.fit",
		false,
		16308592596196249935,
		true,
		tdoNone,
	},
	{
		"sram",
		"Settings.fit",
		false,
		18065441791205275238,
		true,
		tdoNone,
	},
	{
		"sram",
		"Settings2.fit",
		false,
		1732661877833270106,
		true,
		tdoNone,
	},
	{
		"dcrainmaker",
		"Edge810-Vector-2013-08-16-15-35-10.fit",
		false,
		323151606788375988,
		true,
		tdoNone,
	},
	{
		"misc",
		"2013-02-06-12-11-14.fit",
		false,
		11164682768695652056,
		true,
		tdoNone,
	},
	{
		"misc",
		"2015-10-13-08-43-15.fit",
		false,
		873536099792889905,
		true,
		tdoNone,
	},
	{
		"corrupt",
		"activity-filecrc.fit",
		true,
		3999743178100289665,
		true,
		tdoNone,
	},
	{
		"corrupt",
		"activity-unexpected-eof.fit",
		true,
		11827134768800481565,
		true,
		tdoNone,
	},
}
