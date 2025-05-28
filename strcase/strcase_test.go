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

func TestToSnake(t *testing.T) {
	tests := []struct {
		s, exp string
	}{
		{"", ""},
		{"0", ""},
		{"_", ""},
		{"x_", "x"},
		{"_x", "x"},
		{"_x_", "x"},
		{"a_really_long_name", "a_really_long_name"},
		{"a_really__long_name", "a_really_long_name"},
		{"something_id", "something_id"},
		{"something_ids", "something_ids"},
		{"acl", "acl"},
		{"acl_", "acl"},
		{"_acl", "acl"},
		{"_acl_", "acl"},
		{"_a_c_l_", "a_c_l"},
		{"gpu_info", "gpu_info"},
		{"gpu_______info", "gpu_info"},
		{"GPU_info", "gpu_info"},
		{"gPU_info", "gpu_info"},
		{"g_p_u_info", "g_p_u_info"},
		{"uuid_id_uuid", "uuid_id_uuid"},
		{"sample_id_ids", "sample_id_ids"},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			if v := ToSnake(test.s); v != test.exp {
				t.Errorf("ToSnake(%q) expected %q, got: %q", test.s, test.exp, v)
			}
		})
	}
}

func TestToKebab(t *testing.T) {
	tests := []struct {
		s, exp string
	}{
		{"", ""},
		{"0", ""},
		{"_", ""},
		{"x_", "x"},
		{"_x", "x"},
		{"_x_", "x"},
		{"a_really_long_name", "a-really-long-name"},
		{"a_really__long_name", "a-really-long-name"},
		{"something_id", "something-id"},
		{"something_ids", "something-ids"},
		{"acl", "acl"},
		{"acl_", "acl"},
		{"_acl", "acl"},
		{"_acl_", "acl"},
		{"_a_c_l_", "a-c-l"},
		{"gpu_info", "gpu-info"},
		{"gpu_______info", "gpu-info"},
		{"GPU_info", "gpu-info"},
		{"gPU_info", "gpu-info"},
		{"g_p_u_info", "g-p-u-info"},
		{"uuid_id_uuid", "uuid-id-uuid"},
		{"sample_id_ids", "sample-id-ids"},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			if v := ToKebab(test.s); v != test.exp {
				t.Errorf("ToKebab(%q) expected %q, got: %q", test.s, test.exp, v)
			}
		})
	}
}
