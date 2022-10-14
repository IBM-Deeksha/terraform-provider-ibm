// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink

const (
	dlGatewaysVirtualConnections   = "gateway_vcs"
	dlVCNetworkAccount             = "network_account"
	dlVCNetworkId                  = "network_id"
	dlVCName                       = "name"
	dlVCType                       = "type"
	dlVCCreatedAt                  = "created_at"
	dlVCStatus                     = "status"
	dlGatewayId                    = "gateway"
	ID                             = "id"
	dlVirtualConnectionId          = "virtual_connection_id"
	dlVirtualConnectionName        = "virtual_connection_name"
	dlVirtualConnectionType        = "virtual_connection_type"
	dlActive                       = "active"
	dlAsPrepends                   = "as_prepends"
	dlAuthenticationKey            = "authentication_key"
	dlBfdInterval                  = "bfd_interval"
	dlBfdMultiplier                = "bfd_multiplier"
	dlBfdStatus                    = "bfd_status"
	dlBfdStatusUpdatedAt           = "bfd_status_updated_at"
	dlBgpAsn                       = "bgp_asn"
	dlBgpBaseCidr                  = "bgp_base_cidr"
	dlBgpCerCidr                   = "bgp_cer_cidr"
	dlBgpIbmAsn                    = "bgp_ibm_asn"
	dlBgpIbmCidr                   = "bgp_ibm_cidr"
	dlBgpStatus                    = "bgp_status"
	dlCarrierName                  = "carrier_name"
	dlChangeRequest                = "change_request"
	dlCipherSuite                  = "cipher_suite"
	dlCompletionNoticeRejectReason = "completion_notice_reject_reason"
	dlConfidentialityOffset        = "confidentiality_offset"
	dlGatewayProvisioning          = "configuring"
	dlConnectionMode               = "connection_mode"
	dlCreatedAt                    = "created_at"
	dlGatewayProvisioningRejected  = "create_rejected"
	dlCrossConnectRouter           = "cross_connect_router"
	dlCrn                          = "crn"
	dlCryptographicAlgorithm       = "cryptographic_algorithm"
	dlCustomerName                 = "customer_name"
	dlFallbackCak                  = "fallback_cak"
	dlGlobal                       = "global"
	dlKeyServerPriority            = "key_server_priority"
	dlLength                       = "length"
	dlLoaRejectReason              = "loa_reject_reason"
	dlLocationDisplayName          = "location_display_name"
	dlLocationName                 = "location_name"
	dlLinkStatus                   = "link_status"
	dlMacSecConfig                 = "macsec_config"
	dlMetered                      = "metered"
	dlName                         = "name"
	dlOperationalStatus            = "operational_status"
	dlPolicy                       = "policy"
	dlPort                         = "port"
	dlPrimaryCak                   = "primary_cak"
	dlProviderAPIManaged           = "provider_api_managed"
	dlGatewayProvisioningDone      = "provisioned"
	dlResourceGroup                = "resource_group"
	dlSakExpiryTime                = "sak_expiry_time"
	dlSpeedMbps                    = "speed_mbps"
	dlMacSecConfigStatus           = "status"
	dlTags                         = "tags"
	dlType                         = "type"
	dlUpdatedAt                    = "updated_at"
	dlVlan                         = "vlan"
	dlWindowSize                   = "window_size"
	customerAccountID              = "customer_account_id"
	dlRouteReports                 = "route_reports"
	dlPrefix                       = "prefix"
	dlRouteReportNextHop           = "next_hop"
	dlGatewayRoutes                = "gateway_routes"
	dlOnPremRoutes                 = "on_prem_routes"
	dlOverlappingRoutes            = "overlapping_routes"
	dlRoutes                       = "routes"
	dlRouteReportStatus            = "status"
	dlVirtualConnectionRoutes      = "virtual_connection_routes"
	dlId                           = "id"
	dlRouteReportPending           = "pending"
	dlRouteReportComplete          = "complete"
	dlRouteReportId                = "route_report_id"
	dlResourceId                   = "id"
)

func NewInt64Pointer(v int64) *int64 {
	return &v
}

func NewStrPointer(v string) *string {
	return &v
}
