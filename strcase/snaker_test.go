package strcase

import "testing"

func TestCamelToSnake(t *testing.T) {
	tests := []struct {
		s, exp string
	}{
		{"", ""},
		{"0", "0"},
		{"_", "_"},
		{"-X-", "-x-"},
		{"-X_", "-x_"},
		{"AReallyLongName", "a_really_long_name"},
		{"SomethingID", "something_id"},
		{"SomethingID_", "something_id_"},
		{"_SomethingID_", "_something_id_"},
		{"_Something-ID_", "_something-id_"},
		{"_Something-IDS_", "_something-ids_"},
		{"_Something-IDs_", "_something-ids_"},
		{"ACL", "acl"},
		{"GPU", "gpu"},
		{"zGPU", "z_gpu"},
		{"GPUs", "gpus"},
		{"!GPU*", "!gpu*"},
		{"GpuInfo", "gpu_info"},
		{"GPUInfo", "gpu_info"},
		{"gpUInfo", "gp_ui_nfo"},
		{"gpUIDNfo", "gp_uid_nfo"},
		{"gpUIDnfo", "gp_uid_nfo"},
		{"HTTPWriter", "http_writer"},
		{"uHTTPWriter", "u_http_writer"},
		{"UHTTPWriter", "u_h_t_t_p_writer"},
		{"UHTTP_Writer", "u_h_t_t_p_writer"},
		{"UHTTP-Writer", "u_h_t_t_p-writer"},
		{"HTTPHTTP", "http_http"},
		{"uHTTPHTTP", "u_http_http"},
		{"uHTTPHTTPS", "u_http_https"},
		{"uHTTPHTTPS*", "u_http_https*"},
		{"uHTTPSUID*", "u_https_uid*"},
		{"UIDuuidUIDIDUUID", "uid_uuid_uid_id_uuid"},
		{"UID-uuidUIDIDUUID", "uid-uuid_uid_id_uuid"},
		{"UIDzuuidUIDIDUUID", "uid_zuuid_uid_id_uuid"},
		{"UIDzUUIDUIDidUUID", "uid_z_uuid_uid_id_uuid"},
		{"UIDzUUID-UIDidUUID", "uid_z_uuid-uid_id_uuid"},
		{"sampleIDIDS", "sample_id_ids"},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			if v := CamelToSnake(test.s); v != test.exp {
				t.Errorf("%q expected %q, got: %q", test.s, test.exp, v)
			}
		})
	}
}

func TestCamelToSnakeIdentifier(t *testing.T) {
	tests := []struct {
		s, exp string
	}{
		{"", ""},
		{"0", ""},
		{"_", ""},
		{"-X-", "x"},
		{"-X_", "x"},
		{"AReallyLongName", "a_really_long_name"},
		{"SomethingID", "something_id"},
		{"SomethingID_", "something_id"},
		{"_SomethingID_", "something_id"},
		{"_Something-ID_", "something_id"},
		{"_Something-IDS_", "something_ids"},
		{"_Something-IDs_", "something_ids"},
		{"ACL", "acl"},
		{"GPU", "gpu"},
		{"zGPU", "z_gpu"},
		{"!GPU*", "gpu"},
		{"GpuInfo", "gpu_info"},
		{"GPUInfo", "gpu_info"},
		{"gpUInfo", "gp_ui_nfo"},
		{"gpUIDNfo", "gp_uid_nfo"},
		{"gpUIDnfo", "gp_uid_nfo"},
		{"HTTPWriter", "http_writer"},
		{"uHTTPWriter", "u_http_writer"},
		{"UHTTPWriter", "u_h_t_t_p_writer"},
		{"UHTTP_Writer", "u_h_t_t_p_writer"},
		{"UHTTP-Writer", "u_h_t_t_p_writer"},
		{"HTTPHTTP", "http_http"},
		{"uHTTPHTTP", "u_http_http"},
		{"uHTTPHTTPS", "u_http_https"},
		{"uHTTPHTTPS*", "u_http_https"},
		{"uHTTPSUID*", "u_https_uid"},
		{"UIDuuidUIDIDUUID", "uid_uuid_uid_id_uuid"},
		{"UID-uuidUIDIDUUID", "uid_uuid_uid_id_uuid"},
		{"UIDzuuidUIDIDUUID", "uid_zuuid_uid_id_uuid"},
		{"UIDzUUIDUIDidUUID", "uid_z_uuid_uid_id_uuid"},
		{"UIDzUUID-UIDidUUID", "uid_z_uuid_uid_id_uuid"},
		{"SampleIDs", "sample_ids"},
		{"SampleIDS", "sample_ids"},
		{"SampleIDIDs", "sample_id_ids"},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			if v := CamelToSnakeIdentifier(test.s); v != test.exp {
				t.Errorf("CamelToSnake(%q) expected %q, got: %q", test.s, test.exp, v)
			}
		})
	}
}

func TestSnakeToCamel(t *testing.T) {
	tests := []struct {
		s, exp string
	}{
		{"", ""},
		{"0", "0"},
		{"_", ""},
		{"x_", "X"},
		{"_x", "X"},
		{"_x_", "X"},
		{"a_really_long_name", "AReallyLongName"},
		{"a_really__long_name", "AReallyLongName"},
		{"something_id", "SomethingID"},
		{"something_ids", "SomethingIDs"},
		{"acl", "ACL"},
		{"acl_", "ACL"},
		{"_acl", "ACL"},
		{"_acl_", "ACL"},
		{"_a_c_l_", "ACL"},
		{"gpu_info", "GPUInfo"},
		{"gpu_______info", "GPUInfo"},
		{"GPU_info", "GPUInfo"},
		{"gPU_info", "GPUInfo"},
		{"g_p_u_info", "GPUInfo"},
		{"uuid_id_uuid", "UUIDIDUUID"},
		{"sample_id_ids", "SampleIDIDs"},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			if v := SnakeToCamel(test.s); v != test.exp {
				t.Errorf("SnakeToCamel(%q) expected %q, got: %q", test.s, test.exp, v)
			}
		})
	}
}

func TestSnakeToCamelIdentifier(t *testing.T) {
	tests := []struct {
		s, exp string
	}{
		{"", ""},
		{"_", ""},
		{"0", ""},
		{"000", ""},
		{"_000", ""},
		{"_000", ""},
		{"000_", ""},
		{"_000_", ""},
		{"___0--00_", ""},
		{"A0", "A0"},
		{"a_0", "A0"},
		{"a-0", "A0"},
		{"x_", "X"},
		{"_x", "X"},
		{"_x_", "X"},
		{"a_really_long_name", "AReallyLongName"},
		{"_a_really_long_name", "AReallyLongName"},
		{"a_really_long_name_", "AReallyLongName"},
		{"_a_really_long_name_", "AReallyLongName"},
		{"_a_really___long_name_", "AReallyLongName"},
		{"something_id", "SomethingID"},
		{"something-id", "SomethingID"},
		{"-something-id", "SomethingID"},
		{"something-id-", "SomethingID"},
		{"-something-id-", "SomethingID"},
		{"-something_ids-", "SomethingIDs"},
		{"-something_id_s-", "SomethingIDS"},
		{"g_p_u_s", "GPUS"},
		{"gpus", "GPUs"},
		{"acl", "ACL"},
		{"acls", "ACLs"},
		{"acl_", "ACL"},
		{"_acl", "ACL"},
		{"_acl_", "ACL"},
		{"_a_c_l_", "ACL"},
		{"gpu_info", "GPUInfo"},
		{"g_p_u_info", "GPUInfo"},
		{"uuid_id_uuid", "UUIDIDUUID"},
		{"sample_id_ids", "SampleIDIDs"},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			if v := SnakeToCamelIdentifier(test.s); v != test.exp {
				t.Errorf("SnakeToCamelIdentifier(%q) expected %q, got: %q", test.s, test.exp, v)
			}
		})
	}
}
