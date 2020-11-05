package apk

import (
	"github.com/shogo82148/androidbinary"
)

// Instrumentation is an application instrumentation code.
type Instrumentation struct {
	Name            androidbinary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Target          androidbinary.String `xml:"http://schemas.android.com/apk/res/android targetPackage,attr"`
	HandleProfiling androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android handleProfiling,attr"`
	FunctionalTest  androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android functionalTest,attr"`
}

// ActivityAction is an action of an activity.
type ActivityAction struct {
	Name androidbinary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
}

// ActivityCategory is a category of an activity.
type ActivityCategory struct {
	Name androidbinary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
}

// ActivityIntentFilter is an androidbinary.Int32ent filter of an activity.
type ActivityIntentFilter struct {
	Actions    []ActivityAction   `xml:"action"`
	Categories []ActivityCategory `xml:"category"`
}

// AppActivity is an activity in an application.
type AppActivity struct {
	Theme             androidbinary.String   `xml:"http://schemas.android.com/apk/res/android theme,attr"`
	Name              androidbinary.String   `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Label             androidbinary.String   `xml:"http://schemas.android.com/apk/res/android label,attr"`
	ScreenOrientation androidbinary.String   `xml:"http://schemas.android.com/apk/res/android screenOrientation,attr"`
	IntentFilters     []ActivityIntentFilter `xml:"intent-filter"`
	MetaData          []MetaData             `xml:"meta-data"`
}

// AppActivityAlias https://developer.android.com/guide/topics/manifest/activity-alias-element
type AppActivityAlias struct {
	Name           androidbinary.String   `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Label          androidbinary.String   `xml:"http://schemas.android.com/apk/res/android label,attr"`
	TargetActivity androidbinary.String   `xml:"http://schemas.android.com/apk/res/android targetActivity,attr"`
	IntentFilters  []ActivityIntentFilter `xml:"intent-filter"`
	MetaData       []MetaData             `xml:"meta-data"`
}

// MetaData is a metadata
// contained in <activity> <activity-alias> <application> <provider> <receiver> <service>
type MetaData struct {
	Name  androidbinary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Value androidbinary.String `xml:"http://schemas.android.com/apk/res/android value,attr"`
}

// Provider is a provider in an application.
type Provider struct {
	Authorities         androidbinary.String `xml:"http://schemas.android.com/apk/res/android authorities,attr"`
	DirectBootAware     androidbinary.String `xml:"http://schemas.android.com/apk/res/android directBootAware,attr"`
	Enable              androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android enable,attr"`
	Exported            androidbinary.String `xml:"http://schemas.android.com/apk/res/android exported,attr"`
	GrantUriPermissions androidbinary.String `xml:"http://schemas.android.com/apk/res/android grantUriPermissions,attr"`
	Icon                androidbinary.String `xml:"http://schemas.android.com/apk/res/android icon,attr"`
	InitOrder           androidbinary.String `xml:"http://schemas.android.com/apk/res/android initOrder,attr"`
	Label               androidbinary.String `xml:"http://schemas.android.com/apk/res/android label,attr"`
	Multiprocess        androidbinary.String `xml:"http://schemas.android.com/apk/res/android multiprocess,attr"`
	Name                androidbinary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Permission          androidbinary.String `xml:"http://schemas.android.com/apk/res/android permission,attr"`
	Process             androidbinary.String `xml:"http://schemas.android.com/apk/res/android process,attr"`
	ReadPermission      androidbinary.String `xml:"http://schemas.android.com/apk/res/android readPermission,attr"`
	Syncable            androidbinary.String `xml:"http://schemas.android.com/apk/res/android syncable,attr"`
	WritePermission     androidbinary.String `xml:"http://schemas.android.com/apk/res/android writePermission,attr"`
	MetaData            []MetaData           `xml:"meta-data"`
}

// Application is an application in an APK.
type Application struct {
	AllowTaskReparenting  androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android allowTaskReparenting,attr"`
	AllowBackup           androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android allowBackup,attr"`
	AllowClearUserData    androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android allowClearUserData,attr"`
	BackupAgent           androidbinary.String `xml:"http://schemas.android.com/apk/res/android backupAgent,attr"`
	BackupInForeground    androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android backupInForeground,attr"`
	Banner                androidbinary.String `xml:"http://schemas.android.com/apk/res/android banner,attr"`
	Debuggable            androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android debuggable,attr"`
	Description           androidbinary.String `xml:"http://schemas.android.com/apk/res/android description,attr"`
	DirectBootAware       androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android directBootAware,attr"`
	Enabled               androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android enabled,attr"`
	ExtractNativeLibs     androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android extractNativeLibs,attr"`
	FullBackupContent     androidbinary.String `xml:"http://schemas.android.com/apk/res/android fullBackupContent,attr"`
	FullBackupOnly        androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android fullBackupOnly,attr"`
	HasCode               androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android hasCode,attr"`
	HardwareAccelerated   androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android hardwareAccelerated,attr"`
	Icon                  androidbinary.String `xml:"http://schemas.android.com/apk/res/android icon,attr"`
	IsGame                androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android isGame,attr"`
	KillAfterRestore      androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android killAfterRestore,attr"`
	LargeHeap             androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android largeHeap,attr"`
	Label                 androidbinary.String `xml:"http://schemas.android.com/apk/res/android label,attr"`
	Logo                  androidbinary.String `xml:"http://schemas.android.com/apk/res/android logo,attr"`
	ManageSpaceActivity   androidbinary.String `xml:"http://schemas.android.com/apk/res/android manageSpaceActivity,attr"`
	Name                  androidbinary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
	NetworkSecurityConfig androidbinary.String `xml:"http://schemas.android.com/apk/res/android networkSecurityConfig,attr"`
	Permission            androidbinary.String `xml:"http://schemas.android.com/apk/res/android permission,attr"`
	Persistent            androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android persistent,attr"`
	Process               androidbinary.String `xml:"http://schemas.android.com/apk/res/android process,attr"`
	RestoreAnyVersion     androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android restoreAnyVersion,attr"`
	RequiredAccountType   androidbinary.String `xml:"http://schemas.android.com/apk/res/android requiredAccountType,attr"`
	ResizeableActivity    androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android resizeableActivity,attr"`
	RestrictedAccountType androidbinary.String `xml:"http://schemas.android.com/apk/res/android restrictedAccountType,attr"`
	SupportsRtl           androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android supportsRtl,attr"`
	TaskAffinity          androidbinary.String `xml:"http://schemas.android.com/apk/res/android taskAffinity,attr"`
	TestOnly              androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android testOnly,attr"`
	Theme                 androidbinary.String `xml:"http://schemas.android.com/apk/res/android theme,attr"`
	UIOptions             androidbinary.String `xml:"http://schemas.android.com/apk/res/android uiOptions,attr"`
	UsesCleartextTraffic  androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android usesCleartextTraffic,attr"`
	VMSafeMode            androidbinary.Bool   `xml:"http://schemas.android.com/apk/res/android vmSafeMode,attr"`
	Activities            []AppActivity        `xml:"activity"`
	Providers             []Provider           `xml:"provider"`
	ActivityAliases       []AppActivityAlias   `xml:"activity-alias"`
	MetaData              []MetaData           `xml:"meta-data"`
}

// UsesSDK is target SDK version.
type UsesSDK struct {
	Min    androidbinary.Int32 `xml:"http://schemas.android.com/apk/res/android minSdkVersion,attr"`
	Target androidbinary.Int32 `xml:"http://schemas.android.com/apk/res/android targetSdkVersion,attr"`
	Max    androidbinary.Int32 `xml:"http://schemas.android.com/apk/res/android maxSdkVersion,attr"`
}

type UsesPermission struct {
	Name          androidbinary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
	MaxSdkVersion androidbinary.Int32  `xml:"http://schemas.android.com/apk/res/android maxSdkVersion,attr"`
}

// Manifest is a manifest of an APK.
type Manifest struct {
	Package     androidbinary.String `xml:"package,attr"`
	VersionCode androidbinary.Int32  `xml:"http://schemas.android.com/apk/res/android versionCode,attr"`
	VersionName androidbinary.String `xml:"http://schemas.android.com/apk/res/android versionName,attr"`
	App         Application          `xml:"application"`
	Instrument  Instrumentation      `xml:"instrumentation"`
	SDK         UsesSDK              `xml:"uses-sdk"`
	Permissions []UsesPermission     `xml:"uses-permission"`
}
