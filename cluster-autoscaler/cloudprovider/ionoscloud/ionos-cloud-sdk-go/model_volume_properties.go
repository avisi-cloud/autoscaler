/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// VolumeProperties struct for VolumeProperties
type VolumeProperties struct {
	// The availability zone in which the volume should be provisioned. The storage volume will be provisioned on as few physical storage devices as possible, but this cannot be guaranteed upfront. This is uavailable for DAS (Direct Attached Storage), and subject to availability for SSD.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	// The ID of the backup unit that the user has access to. The property is immutable and is only allowed to be set on creation of a new a volume. It is mandatory to provide either 'public image' or 'imageAlias' in conjunction with this property.
	BackupunitId *string `json:"backupunitId,omitempty"`
	// Determines whether the volume will be used as a boot volume. Set to `NONE`, the volume will not be used as boot volume. Set to `PRIMARY`, the volume will be used as boot volume and all other volumes must be set to `NONE`. Set to `AUTO` or `null` requires all volumes to be set to `AUTO` or `null`; this will use the legacy behavior, which is to use the volume as a boot volume only if there are no other volumes or cdrom devices.
	// to set this field to `nil` in order to be marshalled, the explicit nil address `Nilstring` can be used, or the setter `SetBootOrderNil`
	BootOrder *string `json:"bootOrder,omitempty"`
	// The UUID of the attached server.
	BootServer *string `json:"bootServer,omitempty"`
	// The bus type for this volume; default is VIRTIO.
	Bus *string `json:"bus,omitempty"`
	// Hot-plug capable CPU (no reboot required).
	CpuHotPlug *bool `json:"cpuHotPlug,omitempty"`
	// The Logical Unit Number of the storage volume. Null for volumes, not mounted to a VM.
	DeviceNumber *int64 `json:"deviceNumber,omitempty"`
	// Hot-plug capable Virt-IO drive (no reboot required).
	DiscVirtioHotPlug *bool `json:"discVirtioHotPlug,omitempty"`
	// Hot-unplug capable Virt-IO drive (no reboot required). Not supported with Windows VMs.
	DiscVirtioHotUnplug *bool `json:"discVirtioHotUnplug,omitempty"`
	// Image or snapshot ID to be used as template for this volume.
	Image      *string `json:"image,omitempty"`
	ImageAlias *string `json:"imageAlias,omitempty"`
	// Initial password to be set for installed OS. Works with public images only. Not modifiable, forbidden in update requests. Password rules allows all characters from a-z, A-Z, 0-9.
	ImagePassword *string `json:"imagePassword,omitempty"`
	// OS type for this volume.
	LicenceType *string `json:"licenceType,omitempty"`
	// The name of the  resource.
	Name *string `json:"name,omitempty"`
	// Hot-plug capable NIC (no reboot required).
	NicHotPlug *bool `json:"nicHotPlug,omitempty"`
	// Hot-unplug capable NIC (no reboot required).
	NicHotUnplug *bool `json:"nicHotUnplug,omitempty"`
	// The PCI slot number of the storage volume. Null for volumes, not mounted to a VM.
	PciSlot *int32 `json:"pciSlot,omitempty"`
	// Hot-plug capable RAM (no reboot required).
	RamHotPlug *bool `json:"ramHotPlug,omitempty"`
	// The size of the volume in GB.
	Size *float32 `json:"size"`
	// Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation.
	SshKeys *[]string `json:"sshKeys,omitempty"`
	// Hardware type of the volume. DAS (Direct Attached Storage) could be used only in a composite call with a Cube server.
	Type *string `json:"type,omitempty"`
	// The cloud-init configuration for the volume as base64 encoded string. The property is immutable and is only allowed to be set on creation of a new a volume. It is mandatory to provide either 'public image' or 'imageAlias' that has cloud-init compatibility in conjunction with this property.
	UserData *string `json:"userData,omitempty"`
}

// NewVolumeProperties instantiates a new VolumeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeProperties(size float32) *VolumeProperties {
	this := VolumeProperties{}

	var bootOrder = "AUTO"
	this.BootOrder = &bootOrder
	this.Size = &size

	return &this
}

// NewVolumePropertiesWithDefaults instantiates a new VolumeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumePropertiesWithDefaults() *VolumeProperties {
	this := VolumeProperties{}
	var bootOrder = "AUTO"
	this.BootOrder = &bootOrder
	return &this
}

// GetAvailabilityZone returns the AvailabilityZone field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetAvailabilityZone() *string {
	if o == nil {
		return nil
	}

	return o.AvailabilityZone

}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *VolumeProperties) SetAvailabilityZone(v string) {

	o.AvailabilityZone = &v

}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *VolumeProperties) HasAvailabilityZone() bool {
	if o != nil && o.AvailabilityZone != nil {
		return true
	}

	return false
}

// GetBackupunitId returns the BackupunitId field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetBackupunitId() *string {
	if o == nil {
		return nil
	}

	return o.BackupunitId

}

// GetBackupunitIdOk returns a tuple with the BackupunitId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetBackupunitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.BackupunitId, true
}

// SetBackupunitId sets field value
func (o *VolumeProperties) SetBackupunitId(v string) {

	o.BackupunitId = &v

}

// HasBackupunitId returns a boolean if a field has been set.
func (o *VolumeProperties) HasBackupunitId() bool {
	if o != nil && o.BackupunitId != nil {
		return true
	}

	return false
}

// GetBootOrder returns the BootOrder field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetBootOrder() *string {
	if o == nil {
		return nil
	}

	return o.BootOrder

}

// GetBootOrderOk returns a tuple with the BootOrder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetBootOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.BootOrder, true
}

// SetBootOrder sets field value
func (o *VolumeProperties) SetBootOrder(v string) {

	o.BootOrder = &v

}

// sets BootOrder to the explicit address that will be encoded as nil when marshaled
func (o *VolumeProperties) SetBootOrderNil() {
	o.BootOrder = &Nilstring
}

// HasBootOrder returns a boolean if a field has been set.
func (o *VolumeProperties) HasBootOrder() bool {
	if o != nil && o.BootOrder != nil {
		return true
	}

	return false
}

// GetBootServer returns the BootServer field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetBootServer() *string {
	if o == nil {
		return nil
	}

	return o.BootServer

}

// GetBootServerOk returns a tuple with the BootServer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetBootServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.BootServer, true
}

// SetBootServer sets field value
func (o *VolumeProperties) SetBootServer(v string) {

	o.BootServer = &v

}

// HasBootServer returns a boolean if a field has been set.
func (o *VolumeProperties) HasBootServer() bool {
	if o != nil && o.BootServer != nil {
		return true
	}

	return false
}

// GetBus returns the Bus field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetBus() *string {
	if o == nil {
		return nil
	}

	return o.Bus

}

// GetBusOk returns a tuple with the Bus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetBusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Bus, true
}

// SetBus sets field value
func (o *VolumeProperties) SetBus(v string) {

	o.Bus = &v

}

// HasBus returns a boolean if a field has been set.
func (o *VolumeProperties) HasBus() bool {
	if o != nil && o.Bus != nil {
		return true
	}

	return false
}

// GetCpuHotPlug returns the CpuHotPlug field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetCpuHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.CpuHotPlug

}

// GetCpuHotPlugOk returns a tuple with the CpuHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetCpuHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.CpuHotPlug, true
}

// SetCpuHotPlug sets field value
func (o *VolumeProperties) SetCpuHotPlug(v bool) {

	o.CpuHotPlug = &v

}

// HasCpuHotPlug returns a boolean if a field has been set.
func (o *VolumeProperties) HasCpuHotPlug() bool {
	if o != nil && o.CpuHotPlug != nil {
		return true
	}

	return false
}

// GetDeviceNumber returns the DeviceNumber field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetDeviceNumber() *int64 {
	if o == nil {
		return nil
	}

	return o.DeviceNumber

}

// GetDeviceNumberOk returns a tuple with the DeviceNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetDeviceNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}

	return o.DeviceNumber, true
}

// SetDeviceNumber sets field value
func (o *VolumeProperties) SetDeviceNumber(v int64) {

	o.DeviceNumber = &v

}

// HasDeviceNumber returns a boolean if a field has been set.
func (o *VolumeProperties) HasDeviceNumber() bool {
	if o != nil && o.DeviceNumber != nil {
		return true
	}

	return false
}

// GetDiscVirtioHotPlug returns the DiscVirtioHotPlug field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetDiscVirtioHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.DiscVirtioHotPlug

}

// GetDiscVirtioHotPlugOk returns a tuple with the DiscVirtioHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetDiscVirtioHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.DiscVirtioHotPlug, true
}

// SetDiscVirtioHotPlug sets field value
func (o *VolumeProperties) SetDiscVirtioHotPlug(v bool) {

	o.DiscVirtioHotPlug = &v

}

// HasDiscVirtioHotPlug returns a boolean if a field has been set.
func (o *VolumeProperties) HasDiscVirtioHotPlug() bool {
	if o != nil && o.DiscVirtioHotPlug != nil {
		return true
	}

	return false
}

// GetDiscVirtioHotUnplug returns the DiscVirtioHotUnplug field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetDiscVirtioHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.DiscVirtioHotUnplug

}

// GetDiscVirtioHotUnplugOk returns a tuple with the DiscVirtioHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetDiscVirtioHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.DiscVirtioHotUnplug, true
}

// SetDiscVirtioHotUnplug sets field value
func (o *VolumeProperties) SetDiscVirtioHotUnplug(v bool) {

	o.DiscVirtioHotUnplug = &v

}

// HasDiscVirtioHotUnplug returns a boolean if a field has been set.
func (o *VolumeProperties) HasDiscVirtioHotUnplug() bool {
	if o != nil && o.DiscVirtioHotUnplug != nil {
		return true
	}

	return false
}

// GetImage returns the Image field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetImage() *string {
	if o == nil {
		return nil
	}

	return o.Image

}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Image, true
}

// SetImage sets field value
func (o *VolumeProperties) SetImage(v string) {

	o.Image = &v

}

// HasImage returns a boolean if a field has been set.
func (o *VolumeProperties) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// GetImageAlias returns the ImageAlias field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetImageAlias() *string {
	if o == nil {
		return nil
	}

	return o.ImageAlias

}

// GetImageAliasOk returns a tuple with the ImageAlias field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetImageAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ImageAlias, true
}

// SetImageAlias sets field value
func (o *VolumeProperties) SetImageAlias(v string) {

	o.ImageAlias = &v

}

// HasImageAlias returns a boolean if a field has been set.
func (o *VolumeProperties) HasImageAlias() bool {
	if o != nil && o.ImageAlias != nil {
		return true
	}

	return false
}

// GetImagePassword returns the ImagePassword field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetImagePassword() *string {
	if o == nil {
		return nil
	}

	return o.ImagePassword

}

// GetImagePasswordOk returns a tuple with the ImagePassword field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetImagePasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ImagePassword, true
}

// SetImagePassword sets field value
func (o *VolumeProperties) SetImagePassword(v string) {

	o.ImagePassword = &v

}

// HasImagePassword returns a boolean if a field has been set.
func (o *VolumeProperties) HasImagePassword() bool {
	if o != nil && o.ImagePassword != nil {
		return true
	}

	return false
}

// GetLicenceType returns the LicenceType field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetLicenceType() *string {
	if o == nil {
		return nil
	}

	return o.LicenceType

}

// GetLicenceTypeOk returns a tuple with the LicenceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetLicenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.LicenceType, true
}

// SetLicenceType sets field value
func (o *VolumeProperties) SetLicenceType(v string) {

	o.LicenceType = &v

}

// HasLicenceType returns a boolean if a field has been set.
func (o *VolumeProperties) HasLicenceType() bool {
	if o != nil && o.LicenceType != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *VolumeProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *VolumeProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetNicHotPlug returns the NicHotPlug field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetNicHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.NicHotPlug

}

// GetNicHotPlugOk returns a tuple with the NicHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetNicHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.NicHotPlug, true
}

// SetNicHotPlug sets field value
func (o *VolumeProperties) SetNicHotPlug(v bool) {

	o.NicHotPlug = &v

}

// HasNicHotPlug returns a boolean if a field has been set.
func (o *VolumeProperties) HasNicHotPlug() bool {
	if o != nil && o.NicHotPlug != nil {
		return true
	}

	return false
}

// GetNicHotUnplug returns the NicHotUnplug field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetNicHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.NicHotUnplug

}

// GetNicHotUnplugOk returns a tuple with the NicHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetNicHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.NicHotUnplug, true
}

// SetNicHotUnplug sets field value
func (o *VolumeProperties) SetNicHotUnplug(v bool) {

	o.NicHotUnplug = &v

}

// HasNicHotUnplug returns a boolean if a field has been set.
func (o *VolumeProperties) HasNicHotUnplug() bool {
	if o != nil && o.NicHotUnplug != nil {
		return true
	}

	return false
}

// GetPciSlot returns the PciSlot field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetPciSlot() *int32 {
	if o == nil {
		return nil
	}

	return o.PciSlot

}

// GetPciSlotOk returns a tuple with the PciSlot field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetPciSlotOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.PciSlot, true
}

// SetPciSlot sets field value
func (o *VolumeProperties) SetPciSlot(v int32) {

	o.PciSlot = &v

}

// HasPciSlot returns a boolean if a field has been set.
func (o *VolumeProperties) HasPciSlot() bool {
	if o != nil && o.PciSlot != nil {
		return true
	}

	return false
}

// GetRamHotPlug returns the RamHotPlug field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetRamHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.RamHotPlug

}

// GetRamHotPlugOk returns a tuple with the RamHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetRamHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.RamHotPlug, true
}

// SetRamHotPlug sets field value
func (o *VolumeProperties) SetRamHotPlug(v bool) {

	o.RamHotPlug = &v

}

// HasRamHotPlug returns a boolean if a field has been set.
func (o *VolumeProperties) HasRamHotPlug() bool {
	if o != nil && o.RamHotPlug != nil {
		return true
	}

	return false
}

// GetSize returns the Size field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetSize() *float32 {
	if o == nil {
		return nil
	}

	return o.Size

}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Size, true
}

// SetSize sets field value
func (o *VolumeProperties) SetSize(v float32) {

	o.Size = &v

}

// HasSize returns a boolean if a field has been set.
func (o *VolumeProperties) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// GetSshKeys returns the SshKeys field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetSshKeys() *[]string {
	if o == nil {
		return nil
	}

	return o.SshKeys

}

// GetSshKeysOk returns a tuple with the SshKeys field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetSshKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.SshKeys, true
}

// SetSshKeys sets field value
func (o *VolumeProperties) SetSshKeys(v []string) {

	o.SshKeys = &v

}

// HasSshKeys returns a boolean if a field has been set.
func (o *VolumeProperties) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *VolumeProperties) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *VolumeProperties) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetUserData returns the UserData field value
// If the value is explicit nil, nil is returned
func (o *VolumeProperties) GetUserData() *string {
	if o == nil {
		return nil
	}

	return o.UserData

}

// GetUserDataOk returns a tuple with the UserData field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VolumeProperties) GetUserDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.UserData, true
}

// SetUserData sets field value
func (o *VolumeProperties) SetUserData(v string) {

	o.UserData = &v

}

// HasUserData returns a boolean if a field has been set.
func (o *VolumeProperties) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

func (o VolumeProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvailabilityZone != nil {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}

	if o.BackupunitId != nil {
		toSerialize["backupunitId"] = o.BackupunitId
	}

	if o.BootOrder == &Nilstring {
		toSerialize["bootOrder"] = nil
	} else if o.BootOrder != nil {
		toSerialize["bootOrder"] = o.BootOrder
	}
	if o.BootServer != nil {
		toSerialize["bootServer"] = o.BootServer
	}

	if o.Bus != nil {
		toSerialize["bus"] = o.Bus
	}

	if o.CpuHotPlug != nil {
		toSerialize["cpuHotPlug"] = o.CpuHotPlug
	}

	if o.DeviceNumber != nil {
		toSerialize["deviceNumber"] = o.DeviceNumber
	}

	if o.DiscVirtioHotPlug != nil {
		toSerialize["discVirtioHotPlug"] = o.DiscVirtioHotPlug
	}

	if o.DiscVirtioHotUnplug != nil {
		toSerialize["discVirtioHotUnplug"] = o.DiscVirtioHotUnplug
	}

	if o.Image != nil {
		toSerialize["image"] = o.Image
	}

	if o.ImageAlias != nil {
		toSerialize["imageAlias"] = o.ImageAlias
	}

	if o.ImagePassword != nil {
		toSerialize["imagePassword"] = o.ImagePassword
	}

	if o.LicenceType != nil {
		toSerialize["licenceType"] = o.LicenceType
	}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.NicHotPlug != nil {
		toSerialize["nicHotPlug"] = o.NicHotPlug
	}

	if o.NicHotUnplug != nil {
		toSerialize["nicHotUnplug"] = o.NicHotUnplug
	}

	if o.PciSlot != nil {
		toSerialize["pciSlot"] = o.PciSlot
	}

	if o.RamHotPlug != nil {
		toSerialize["ramHotPlug"] = o.RamHotPlug
	}

	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	if o.SshKeys != nil {
		toSerialize["sshKeys"] = o.SshKeys
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.UserData != nil {
		toSerialize["userData"] = o.UserData
	}

	return json.Marshal(toSerialize)
}

type NullableVolumeProperties struct {
	value *VolumeProperties
	isSet bool
}

func (v NullableVolumeProperties) Get() *VolumeProperties {
	return v.value
}

func (v *NullableVolumeProperties) Set(val *VolumeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeProperties(val *VolumeProperties) *NullableVolumeProperties {
	return &NullableVolumeProperties{value: val, isSet: true}
}

func (v NullableVolumeProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
