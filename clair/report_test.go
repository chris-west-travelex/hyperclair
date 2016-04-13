package clair

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestReportAsHtml(t *testing.T) {
	var analysis ImageAnalysis
	err := json.Unmarshal([]byte(sampleJSON), &analysis)

	if err != nil {
		t.Errorf("Failing with error: %v", err)
	}

	html, err := ReportAsHTML(analysis)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(os.TempDir()+"/hyperclair-html-report.html", []byte(html), 0700)
	if err != nil {
		log.Fatal(err)
	}
}

const sampleJSON = `
{
	"Registry": "registry:5000",
	"ImageName": "jgsqware/ubuntu-git",
	"Tag": "latest",
	"Layers": [
		{
			"Layer": {
				"Name": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845",
				"Namespace": "ubuntu:14.04",
				"ParentName": "sha256:9e0bc8a71bde464f710bc2b593a1fc21521517671e918687892303151331fa56",
				"IndexedByVersion": 2,
				"Features": [
					{
						"Name": "sudo",
						"Namespace": "ubuntu:14.04",
						"Version": "1.8.9p5-1ubuntu1.2",
						"Vulnerabilities": [
							{
								"Name": "CVE-2015-8239",
								"Namespace": "ubuntu:14.04",
								"Description": "race condition checking digests/checksums in sudoers",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8239",
								"Severity": "Medium"
							},
							{
								"Name": "CVE-2015-5602",
								"Namespace": "ubuntu:14.04",
								"Description": "sudoedit in Sudo before 1.8.15 allows local users to gain privileges via a symlink attack on a file whose full path is defined using multiple wildcards in /etc/sudoers, as demonstrated by \"/home/*/*/file.txt.\"",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-5602",
								"Severity": "Medium"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libtext-soundex-perl",
						"Namespace": "ubuntu:14.04",
						"Version": "3.4-1build1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "dash",
						"Namespace": "ubuntu:14.04",
						"Version": "0.5.7-4ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "eject",
						"Namespace": "ubuntu:14.04",
						"Version": "2.1.5+deb1+cvs20081104-13.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "gcc-4.8",
						"Namespace": "ubuntu:14.04",
						"Version": "4.8.4-2ubuntu1~14.04",
						"Vulnerabilities": [
							{
								"Name": "CVE-2014-5044",
								"Namespace": "ubuntu:14.04",
								"Description": "Array memory allocations could cause an integer overflow and thus memory overflow issues at runtime.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-5044",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-5276",
								"Namespace": "ubuntu:14.04",
								"Description": "The std::random_device class in libstdc++ in the GNU Compiler Collection (aka GCC) before 4.9.4 does not properly handle short reads from blocking sources, which makes it easier for context-dependent attackers to predict the random values via unspecified vectors.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-5276",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "lvm2",
						"Namespace": "ubuntu:14.04",
						"Version": "2.02.98-6ubuntu2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "net-tools",
						"Namespace": "ubuntu:14.04",
						"Version": "1.60-25ubuntu2.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libsepol",
						"Namespace": "ubuntu:14.04",
						"Version": "2.2-1ubuntu0.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libtext-charwidth-perl",
						"Namespace": "ubuntu:14.04",
						"Version": "0.04-7build3",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "cdebconf",
						"Namespace": "ubuntu:14.04",
						"Version": "0.187ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "vim",
						"Namespace": "ubuntu:14.04",
						"Version": "2:7.4.052-1ubuntu3",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "attr",
						"Namespace": "ubuntu:14.04",
						"Version": "1:2.4.47-1ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "gnutls26",
						"Namespace": "ubuntu:14.04",
						"Version": "2.12.23-12ubuntu2.3",
						"Vulnerabilities": [
							{
								"Name": "CVE-2015-7575",
								"Namespace": "ubuntu:14.04",
								"Description": "Mozilla Network Security Services (NSS) before 3.20.2, as used in Mozilla Firefox before 43.0.2 and Firefox ESR 38.x before 38.5.2, does not reject MD5 signatures in Server Key Exchange messages in TLS 1.2 Handshake Protocol traffic, which makes it easier for man-in-the-middle attackers to spoof servers by triggering a collision.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-7575",
								"Severity": "Medium",
								"FixedBy": "2.12.23-12ubuntu2.4"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "python3.4",
						"Namespace": "ubuntu:14.04",
						"Version": "3.4.3-1ubuntu1~14.04.3",
						"Vulnerabilities": [
							{
								"Name": "CVE-2014-2667",
								"Namespace": "ubuntu:14.04",
								"Description": "Race condition in the _get_masked_mode function in Lib/os.py in Python 3.2 through 3.5, when exist_ok is set to true and multiple threads are used, might allow local users to bypass intended file permissions by leveraging a separate application vulnerability before the umask has been set to the expected value.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-2667",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "logrotate",
						"Namespace": "ubuntu:14.04",
						"Version": "3.8.7-1ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "adduser",
						"Namespace": "ubuntu:14.04",
						"Version": "3.113+nmu3ubuntu3",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "base-passwd",
						"Namespace": "ubuntu:14.04",
						"Version": "3.5.33",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "grep",
						"Namespace": "ubuntu:14.04",
						"Version": "2.16-1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "dpkg",
						"Namespace": "ubuntu:14.04",
						"Version": "1.17.5ubuntu5.5",
						"Vulnerabilities": [
							{
								"Name": "CVE-2014-8625",
								"Namespace": "ubuntu:14.04",
								"Description": "Multiple format string vulnerabilities in the parse_error_msg function in parsehelp.c in dpkg before 1.17.22 allow remote attackers to cause a denial of service (crash) and possibly execute arbitrary code via format string specifiers in the (1) package or (2) architecture name.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-8625",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "ustr",
						"Namespace": "ubuntu:14.04",
						"Version": "1.0.4-3ubuntu2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "db5.3",
						"Namespace": "ubuntu:14.04",
						"Version": "5.3.28-3ubuntu3",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libtext-iconv-perl",
						"Namespace": "ubuntu:14.04",
						"Version": "1.7-5build2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "gccgo-4.9",
						"Namespace": "ubuntu:14.04",
						"Version": "4.9.1-0ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "xz-utils",
						"Namespace": "ubuntu:14.04",
						"Version": "5.1.1alpha+20120614-2ubuntu2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "netcat-openbsd",
						"Namespace": "ubuntu:14.04",
						"Version": "1.105-7ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "coreutils",
						"Namespace": "ubuntu:14.04",
						"Version": "8.21-1ubuntu5.3",
						"Vulnerabilities": [
							{
								"Name": "CVE-2015-1865",
								"Namespace": "ubuntu:14.04",
								"Description": "\"time of check to time of use\" race condition fts.c",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-1865",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2016-2781",
								"Namespace": "ubuntu:14.04",
								"Description": "nonpriv session can escape to the parent session by using the TIOCSTI ioctl",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-2781",
								"Severity": "Medium"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "cpio",
						"Namespace": "ubuntu:14.04",
						"Version": "2.11+dfsg-1ubuntu1.1",
						"Vulnerabilities": [
							{
								"Name": "CVE-2016-2037",
								"Namespace": "ubuntu:14.04",
								"Description": "The cpio_safer_name_suffix function in util.c in cpio 2.11 allows remote attackers to cause a denial of service (out-of-bounds write) via a crafted cpio file.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-2037",
								"Severity": "Medium",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 4.3,
											"Vectors": "AV:N/AC:M/Au:N/C:N/I:N"
										}
									}
								},
								"FixedBy": "2.11+dfsg-1ubuntu1.2"
							},
							{
								"Name": "CVE-2015-1197",
								"Namespace": "ubuntu:14.04",
								"Description": "cpio 2.11, when using the --no-absolute-filenames option, allows local users to write to arbitrary files via a symlink attack on a file in an archive.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-1197",
								"Severity": "Low",
								"FixedBy": "2.11+dfsg-1ubuntu1.2"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "tar",
						"Namespace": "ubuntu:14.04",
						"Version": "1.27.1-1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libffi",
						"Namespace": "ubuntu:14.04",
						"Version": "3.1~rc1+r3.0.13-12ubuntu0.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libgcrypt11",
						"Namespace": "ubuntu:14.04",
						"Version": "1.5.3-2ubuntu4.2",
						"Vulnerabilities": [
							{
								"Name": "CVE-2015-7511",
								"Namespace": "ubuntu:14.04",
								"Description": "ECDH Key-Extraction via Low-Bandwidth Electromagnetic Attacks on PCs",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-7511",
								"Severity": "Medium",
								"FixedBy": "1.5.3-2ubuntu4.3"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "findutils",
						"Namespace": "ubuntu:14.04",
						"Version": "4.4.2-7",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "gzip",
						"Namespace": "ubuntu:14.04",
						"Version": "1.6-3ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "klibc",
						"Namespace": "ubuntu:14.04",
						"Version": "2.0.3-0ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "gnupg",
						"Namespace": "ubuntu:14.04",
						"Version": "1.4.16-1ubuntu2.3",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "readline6",
						"Namespace": "ubuntu:14.04",
						"Version": "6.3-4ubuntu2",
						"Vulnerabilities": [
							{
								"Name": "CVE-2014-2524",
								"Namespace": "ubuntu:14.04",
								"Description": "The _rl_tropen function in util.c in GNU readline before 6.3 patch 3 allows local users to create or overwrite arbitrary files via a symlink attack on a /var/tmp/rltrace.[PID] file.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-2524",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "lockfile-progs",
						"Namespace": "ubuntu:14.04",
						"Version": "0.1.17",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "fribidi",
						"Namespace": "ubuntu:14.04",
						"Version": "0.19.6-1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "json-c",
						"Namespace": "ubuntu:14.04",
						"Version": "0.11-3ubuntu1.2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "sysvinit",
						"Namespace": "ubuntu:14.04",
						"Version": "2.88dsf-41ubuntu6.2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "base-files",
						"Namespace": "ubuntu:14.04",
						"Version": "7.2ubuntu5.3",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "ifupdown",
						"Namespace": "ubuntu:14.04",
						"Version": "0.7.47.2ubuntu4.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "resolvconf",
						"Namespace": "ubuntu:14.04",
						"Version": "1.69ubuntu1.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "bash",
						"Namespace": "ubuntu:14.04",
						"Version": "4.3-7ubuntu1.5",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "console-setup",
						"Namespace": "ubuntu:14.04",
						"Version": "1.70ubuntu8",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "pcre3",
						"Namespace": "ubuntu:14.04",
						"Version": "1:8.31-2ubuntu2.1",
						"Vulnerabilities": [
							{
								"Name": "CVE-2015-2328",
								"Namespace": "ubuntu:14.04",
								"Description": "PCRE before 8.36 mishandles the /((?(R)a|(?1)))+/ pattern and related patterns with certain recursion, which allows remote attackers to cause a denial of service (segmentation fault) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-2328",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 7.5,
											"Vectors": "AV:N/AC:L/Au:N/C:P/I:P"
										}
									}
								}
							},
							{
								"Name": "CVE-2016-3191",
								"Namespace": "ubuntu:14.04",
								"Description": "The compile_branch function in pcre_compile.c in PCRE 8.x before 8.39 and pcre2_compile.c in PCRE2 before 10.22 mishandles patterns containing an (*ACCEPT) substring in conjunction with nested parentheses, which allows remote attackers to execute arbitrary code or cause a denial of service (stack-based buffer overflow) via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror, aka ZDI-CAN-3542.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-3191",
								"Severity": "Medium",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 7.5,
											"Vectors": "AV:N/AC:L/Au:N/C:P/I:P"
										}
									}
								}
							},
							{
								"Name": "CVE-2015-8394",
								"Namespace": "ubuntu:14.04",
								"Description": "PCRE before 8.38 mishandles the (?(\u003cdigits\u003e) and (?(R\u003cdigits\u003e) conditions, which allows remote attackers to cause a denial of service (integer overflow) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8394",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 7.5,
											"Vectors": "AV:N/AC:L/Au:N/C:P/I:P"
										}
									}
								}
							},
							{
								"Name": "CVE-2015-8391",
								"Namespace": "ubuntu:14.04",
								"Description": "The pcre_compile function in pcre_compile.c in PCRE before 8.38 mishandles certain [: nesting, which allows remote attackers to cause a denial of service (CPU consumption) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8391",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 9,
											"Vectors": "AV:N/AC:L/Au:N/C:P/I:P"
										}
									}
								}
							},
							{
								"Name": "CVE-2015-8390",
								"Namespace": "ubuntu:14.04",
								"Description": "PCRE before 8.38 mishandles the [: and \\\\ substrings in character classes, which allows remote attackers to cause a denial of service (uninitialized memory read) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8390",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 7.5,
											"Vectors": "AV:N/AC:L/Au:N/C:P/I:P"
										}
									}
								}
							},
							{
								"Name": "CVE-2015-8382",
								"Namespace": "ubuntu:14.04",
								"Description": "The match function in pcre_exec.c in PCRE before 8.37 mishandles the /(?:((abcd))|(((?:(?:(?:(?:abc|(?:abcdef))))b)abcdefghi)abc)|((*ACCEPT)))/ pattern and related patterns involving (*ACCEPT), which allows remote attackers to obtain sensitive information from process memory or cause a denial of service (partially initialized memory and application crash) via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror, aka ZDI-CAN-2547.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8382",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 6.4,
											"Vectors": "AV:N/AC:L/Au:N/C:P/I:N"
										}
									}
								}
							},
							{
								"Name": "CVE-2015-8387",
								"Namespace": "ubuntu:14.04",
								"Description": "PCRE before 8.38 mishandles (?123) subroutine calls and related subroutine calls, which allows remote attackers to cause a denial of service (integer overflow) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8387",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 7.5,
											"Vectors": "AV:N/AC:L/Au:N/C:P/I:P"
										}
									}
								}
							},
							{
								"Name": "CVE-2015-8393",
								"Namespace": "ubuntu:14.04",
								"Description": "pcregrep in PCRE before 8.38 mishandles the -q option for binary files, which might allow remote attackers to obtain sensitive information via a crafted file, as demonstrated by a CGI script that sends stdout data to a client.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8393",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 5,
											"Vectors": "AV:N/AC:L/Au:N/C:P/I:N"
										}
									}
								}
							},
							{
								"Name": "CVE-2015-8386",
								"Namespace": "ubuntu:14.04",
								"Description": "PCRE before 8.38 mishandles the interaction of lookbehind assertions and mutually recursive subpatterns, which allows remote attackers to cause a denial of service (buffer overflow) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8386",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 7.5,
											"Vectors": "AV:N/AC:L/Au:N/C:P/I:P"
										}
									}
								}
							},
							{
								"Name": "CVE-2015-8380",
								"Namespace": "ubuntu:14.04",
								"Description": "The pcre_exec function in pcre_exec.c in PCRE before 8.38 mishandles a // pattern with a \\01 string, which allows remote attackers to cause a denial of service (heap-based buffer overflow) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8380",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 7.5,
											"Vectors": "AV:N/AC:L/Au:N/C:P/I:P"
										}
									}
								}
							},
							{
								"Name": "CVE-2015-8385",
								"Namespace": "ubuntu:14.04",
								"Description": "PCRE before 8.38 mishandles the /(?|(\\k'Pm')|(?'Pm'))/ pattern and related patterns with certain forward references, which allows remote attackers to cause a denial of service (buffer overflow) or possibly have unspecified other impact via a crafted regular expression, as demonstrated by a JavaScript RegExp object encountered by Konqueror.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8385",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 7.5,
											"Vectors": "AV:N/AC:L/Au:N/C:P/I:P"
										}
									}
								}
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libsemanage",
						"Namespace": "ubuntu:14.04",
						"Version": "2.2-1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "gdbm",
						"Namespace": "ubuntu:14.04",
						"Version": "1.8.3-12build1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "xkeyboard-config",
						"Namespace": "ubuntu:14.04",
						"Version": "2.10.1-1ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "ncurses",
						"Namespace": "ubuntu:14.04",
						"Version": "5.9+20140118-1ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "python3-defaults",
						"Namespace": "ubuntu:14.04",
						"Version": "3.4.0-0ubuntu2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libarchive-extract-perl",
						"Namespace": "ubuntu:14.04",
						"Version": "0.70-1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "iproute2",
						"Namespace": "ubuntu:14.04",
						"Version": "3.12.0-2ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "liblocale-gettext-perl",
						"Namespace": "ubuntu:14.04",
						"Version": "1.05-7build3",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "mawk",
						"Namespace": "ubuntu:14.04",
						"Version": "1.3.3-17ubuntu2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "expat",
						"Namespace": "ubuntu:14.04",
						"Version": "2.1.0-4ubuntu1.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libcap2",
						"Namespace": "ubuntu:14.04",
						"Version": "1:2.24-0ubuntu2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "initramfs-tools",
						"Namespace": "ubuntu:14.04",
						"Version": "0.103ubuntu4.2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libestr",
						"Namespace": "ubuntu:14.04",
						"Version": "0.1.9-0ubuntu2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libtext-wrapi18n-perl",
						"Namespace": "ubuntu:14.04",
						"Version": "0.06-7",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "debianutils",
						"Namespace": "ubuntu:14.04",
						"Version": "4.4",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "insserv",
						"Namespace": "ubuntu:14.04",
						"Version": "1.14.0-5ubuntu2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "slang2",
						"Namespace": "ubuntu:14.04",
						"Version": "2.2.4-15ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "systemd",
						"Namespace": "ubuntu:14.04",
						"Version": "204-5ubuntu20.15",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "upstart",
						"Namespace": "ubuntu:14.04",
						"Version": "1.12.1-0ubuntu4.2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "kbd",
						"Namespace": "ubuntu:14.04",
						"Version": "1.15.5-1ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "busybox",
						"Namespace": "ubuntu:14.04",
						"Version": "1:1.21.0-1ubuntu1",
						"Vulnerabilities": [
							{
								"Name": "CVE-2014-9645",
								"Namespace": "ubuntu:14.04",
								"Description": "modprobe wrongly accepts paths as module names",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9645",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2011-5325",
								"Namespace": "ubuntu:14.04",
								"Description": "path traversal vulnerability in busybox tar",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2011-5325",
								"Severity": "Medium"
							},
							{
								"Name": "CVE-2016-2147",
								"Namespace": "ubuntu:14.04",
								"Description": "OOB heap write due to integer underflow",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-2147",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2016-2148",
								"Namespace": "ubuntu:14.04",
								"Description": "heap overflow in OPTION_6RD parsing",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-2148",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "eglibc",
						"Namespace": "ubuntu:14.04",
						"Version": "2.19-0ubuntu6.6",
						"Vulnerabilities": [
							{
								"Name": "CVE-2015-8778",
								"Namespace": "ubuntu:14.04",
								"Description": "hcreate((size_t)-1) should fail with ENOMEM",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8778",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2014-9761",
								"Namespace": "ubuntu:14.04",
								"Description": "nan function unbounded stack allocation",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9761",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-5180",
								"Namespace": "ubuntu:14.04",
								"Description": "DNS resolver NULL pointer dereference with crafted record type",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-5180",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2013-2207",
								"Namespace": "ubuntu:14.04",
								"Description": "pt_chown in GNU C Library (aka glibc or libc6) before 2.18 does not properly check permissions for tty files, which allows local users to change the permission on the files and obtain access to arbitrary pseudo-terminals by leveraging a FUSE file system.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2013-2207",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-8776",
								"Namespace": "ubuntu:14.04",
								"Description": "Passing out of range data to strftime() causes a segfault",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8776",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-8779",
								"Namespace": "ubuntu:14.04",
								"Description": "catopen() Multiple unbounded stack allocations",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8779",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-5277",
								"Namespace": "ubuntu:14.04",
								"Description": "The get_contents function in nss_files/files-XXX.c in the Name Service Switch (NSS) in GNU C Library (aka glibc or libc6) before 2.20 might allow local users to cause a denial of service (heap corruption) or gain privileges via a long line in the NSS files database.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-5277",
								"Severity": "Medium"
							},
							{
								"Name": "CVE-2015-7547",
								"Namespace": "ubuntu:14.04",
								"Description": "Multiple stack-based buffer overflows in the (1) send_dg and (2) send_vc functions in the libresolv library in the GNU C Library (aka glibc or libc6) before 2.23 allow remote attackers to cause a denial of service (crash) or possibly execute arbitrary code via a crafted DNS response that triggers a call to the getaddrinfo function with the AF_UNSPEC or AF_INET6 address family, related to performing \"dual A/AAAA DNS queries\" and the libnss_dns.so.2 NSS module.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-7547",
								"Severity": "High",
								"FixedBy": "2.19-0ubuntu6.7"
							},
							{
								"Name": "CVE-2014-8121",
								"Namespace": "ubuntu:14.04",
								"Description": "DB_LOOKUP in nss_files/files-XXX.c in the Name Service Switch (NSS) in GNU C Library (aka glibc or libc6) 2.21 and earlier does not properly check if a file is open, which allows remote attackers to cause a denial of service (infinite loop) by performing a look-up while the database is iterated over the database, which triggers the file pointer to be reset.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-8121",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-1781",
								"Namespace": "ubuntu:14.04",
								"Description": "Buffer overflow in the gethostbyname_r and other unspecified NSS functions in the GNU C Library (aka glibc or libc6) before 2.22 allows context-dependent attackers to cause a denial of service (crash) or execute arbitrary code via a crafted DNS response, which triggers a call with a misaligned buffer.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-1781",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2016-1234",
								"Namespace": "ubuntu:14.04",
								"Description": "glob: buffer overflow with GLOB_ALTDIRFUNC due to incorrect NAME_MAX limit assumption",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-1234",
								"Severity": "Medium"
							},
							{
								"Name": "CVE-2015-8777",
								"Namespace": "ubuntu:14.04",
								"Description": "The process_envvars function in elf/rtld.c in the GNU C Library (aka glibc or libc6) before 2.23 allows local users to bypass a pointer-guarding protection mechanism via a zero value of the LD_POINTER_GUARD environment variable.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8777",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "hostname",
						"Namespace": "ubuntu:14.04",
						"Version": "3.15ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "bzip2",
						"Namespace": "ubuntu:14.04",
						"Version": "1.0.6-5",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libnih",
						"Namespace": "ubuntu:14.04",
						"Version": "1.0.3-4ubuntu25",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "mpdecimal",
						"Namespace": "ubuntu:14.04",
						"Version": "2.4.0-6",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "openssl",
						"Namespace": "ubuntu:14.04",
						"Version": "1.0.1f-1ubuntu2.16",
						"Vulnerabilities": [
							{
								"Name": "CVE-2016-0797",
								"Namespace": "ubuntu:14.04",
								"Description": "Multiple integer overflows in OpenSSL 1.0.1 before 1.0.1s and 1.0.2 before 1.0.2g allow remote attackers to cause a denial of service (heap memory corruption or NULL pointer dereference) or possibly have unspecified other impact via a long digit string that is mishandled by the (1) BN_dec2bn or (2) BN_hex2bn function, related to crypto/bn/bn.h and crypto/bn/bn_print.c.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-0797",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 5,
											"Vectors": "AV:N/AC:L/Au:N/C:N/I:N"
										}
									}
								},
								"FixedBy": "1.0.1f-1ubuntu2.18"
							},
							{
								"Name": "CVE-2016-2842",
								"Namespace": "ubuntu:14.04",
								"Description": "The doapr_outch function in crypto/bio/b_print.c in OpenSSL 1.0.1 before 1.0.1s and 1.0.2 before 1.0.2g does not verify that a certain memory allocation succeeds, which allows remote attackers to cause a denial of service (out-of-bounds write or memory consumption) or possibly have unspecified other impact via a long string, as demonstrated by a large amount of ASN.1 data, a different vulnerability than CVE-2016-0799.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-2842",
								"Severity": "Medium",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 10,
											"Vectors": "AV:N/AC:L/Au:N/C:C/I:C"
										}
									}
								},
								"FixedBy": "1.0.1f-1ubuntu2.18"
							},
							{
								"Name": "CVE-2016-0702",
								"Namespace": "ubuntu:14.04",
								"Description": "The MOD_EXP_CTIME_COPY_FROM_PREBUF function in crypto/bn/bn_exp.c in OpenSSL 1.0.1 before 1.0.1s and 1.0.2 before 1.0.2g does not properly consider cache-bank access times during modular exponentiation, which makes it easier for local users to discover RSA keys by running a crafted application on the same Intel Sandy Bridge CPU core as a victim and leveraging cache-bank conflicts, aka a \"CacheBleed\" attack.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-0702",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 1.9,
											"Vectors": "AV:L/AC:M/Au:N/C:P/I:N"
										}
									}
								},
								"FixedBy": "1.0.1f-1ubuntu2.18"
							},
							{
								"Name": "CVE-2016-0705",
								"Namespace": "ubuntu:14.04",
								"Description": "Double free vulnerability in the dsa_priv_decode function in crypto/dsa/dsa_ameth.c in OpenSSL 1.0.1 before 1.0.1s and 1.0.2 before 1.0.2g allows remote attackers to cause a denial of service (memory corruption) or possibly have unspecified other impact via a malformed DSA private key.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-0705",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 10,
											"Vectors": "AV:N/AC:L/Au:N/C:C/I:C"
										}
									}
								},
								"FixedBy": "1.0.1f-1ubuntu2.18"
							},
							{
								"Name": "CVE-2016-0798",
								"Namespace": "ubuntu:14.04",
								"Description": "Memory leak in the SRP_VBASE_get_by_user implementation in OpenSSL 1.0.1 before 1.0.1s and 1.0.2 before 1.0.2g allows remote attackers to cause a denial of service (memory consumption) by providing an invalid username in a connection attempt, related to apps/s_server.c and crypto/srp/srp_vfy.c.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-0798",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 7.8,
											"Vectors": "AV:N/AC:L/Au:N/C:N/I:N"
										}
									}
								},
								"FixedBy": "1.0.1f-1ubuntu2.18"
							},
							{
								"Name": "CVE-2016-0799",
								"Namespace": "ubuntu:14.04",
								"Description": "The fmtstr function in crypto/bio/b_print.c in OpenSSL 1.0.1 before 1.0.1s and 1.0.2 before 1.0.2g improperly calculates string lengths, which allows remote attackers to cause a denial of service (overflow and out-of-bounds read) or possibly have unspecified other impact via a long string, as demonstrated by a large amount of ASN.1 data, a different vulnerability than CVE-2016-2842.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-0799",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 10,
											"Vectors": "AV:N/AC:L/Au:N/C:C/I:C"
										}
									}
								},
								"FixedBy": "1.0.1f-1ubuntu2.18"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "diffutils",
						"Namespace": "ubuntu:14.04",
						"Version": "1:3.3-1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "apt",
						"Namespace": "ubuntu:14.04",
						"Version": "1.0.1ubuntu2.10",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "debconf",
						"Namespace": "ubuntu:14.04",
						"Version": "1.5.51ubuntu2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "ubuntu-meta",
						"Namespace": "ubuntu:14.04",
						"Version": "1.325",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "rsyslog",
						"Namespace": "ubuntu:14.04",
						"Version": "7.4.4-1ubuntu2.6",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "dh-python",
						"Namespace": "ubuntu:14.04",
						"Version": "1.20140128-1ubuntu8.2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libpng",
						"Namespace": "ubuntu:14.04",
						"Version": "1.2.50-1ubuntu2.14.04.1",
						"Vulnerabilities": [
							{
								"Name": "CVE-2015-8540",
								"Namespace": "ubuntu:14.04",
								"Description": "underflow read in png_check_keyword in pngwutil.c",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8540",
								"Severity": "Medium",
								"FixedBy": "1.2.50-1ubuntu2.14.04.2"
							},
							{
								"Name": "CVE-2015-8472",
								"Namespace": "ubuntu:14.04",
								"Description": "Buffer overflow in the png_set_PLTE function in libpng before 1.0.65, 1.1.x and 1.2.x before 1.2.55, 1.3.x, 1.4.x before 1.4.18, 1.5.x before 1.5.25, and 1.6.x before 1.6.20 allows remote attackers to cause a denial of service (application crash) or possibly have unspecified other impact via a small bit-depth value in an IHDR (aka image header) chunk in a PNG image. NOTE: this vulnerability exists because of an incomplete fix for CVE-2015-8126.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8472",
								"Severity": "Medium",
								"FixedBy": "1.2.50-1ubuntu2.14.04.2"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libmodule-pluggable-perl",
						"Namespace": "ubuntu:14.04",
						"Version": "5.1-1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "mountall",
						"Namespace": "ubuntu:14.04",
						"Version": "2.53",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "cron",
						"Namespace": "ubuntu:14.04",
						"Version": "3.0pl1-124ubuntu2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "acl",
						"Namespace": "ubuntu:14.04",
						"Version": "2.2.52-1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libdrm",
						"Namespace": "ubuntu:14.04",
						"Version": "2.4.60-2~ubuntu14.04.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "util-linux",
						"Namespace": "ubuntu:14.04",
						"Version": "2.20.1-5.1ubuntu20.7",
						"Vulnerabilities": [
							{
								"Name": "CVE-2014-9114",
								"Namespace": "ubuntu:14.04",
								"Description": "blkid command injection",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9114",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2013-0157",
								"Namespace": "ubuntu:14.04",
								"Description": "(a) mount and (b) umount in util-linux 2.14.1, 2.17.2, and probably other versions allow local users to determine the existence of restricted directories by (1) using the --guess-fstype command-line option or (2) attempting to mount a non-existent device, which generates different error messages depending on whether the directory exists.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2013-0157",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "tzdata",
						"Namespace": "ubuntu:14.04",
						"Version": "2015g-0ubuntu0.14.04",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "less",
						"Namespace": "ubuntu:14.04",
						"Version": "458-2",
						"Vulnerabilities": [
							{
								"Name": "CVE-2014-9488",
								"Namespace": "ubuntu:14.04",
								"Description": "The is_utf8_well_formed function in GNU less before 475 allows remote attackers to have unspecified impact via malformed UTF-8 characters, which triggers an out-of-bounds read.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9488",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "popt",
						"Namespace": "ubuntu:14.04",
						"Version": "1.16-8ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "init-system-helpers",
						"Namespace": "ubuntu:14.04",
						"Version": "1.14",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "mime-support",
						"Namespace": "ubuntu:14.04",
						"Version": "3.54ubuntu1.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "p11-kit",
						"Namespace": "ubuntu:14.04",
						"Version": "0.20.2-2ubuntu2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "cgmanager",
						"Namespace": "ubuntu:14.04",
						"Version": "0.24-0ubuntu7.5",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libselinux",
						"Namespace": "ubuntu:14.04",
						"Version": "2.2.2-1ubuntu0.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "langpack-locales",
						"Namespace": "ubuntu:14.04",
						"Version": "2.13+git20120306-12.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "procps",
						"Namespace": "ubuntu:14.04",
						"Version": "1:3.3.9-1ubuntu2.2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "ureadahead",
						"Namespace": "ubuntu:14.04",
						"Version": "0.100.0-16",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "file",
						"Namespace": "ubuntu:14.04",
						"Version": "1:5.14-2ubuntu3.3",
						"Vulnerabilities": [
							{
								"Name": "CVE-2014-9621",
								"Namespace": "ubuntu:14.04",
								"Description": "The ELF parser in file 5.16 through 5.21 allows remote attackers to cause a denial of service via a long string.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9621",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2014-9620",
								"Namespace": "ubuntu:14.04",
								"Description": "The ELF parser in file 5.08 through 5.21 allows remote attackers to cause a denial of service via a large number of notes.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9620",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2014-9653",
								"Namespace": "ubuntu:14.04",
								"Description": "readelf.c in file before 5.22, as used in the Fileinfo component in PHP before 5.4.37, 5.5.x before 5.5.21, and 5.6.x before 5.6.5, does not consider that pread calls sometimes read only a subset of the available data, which allows remote attackers to cause a denial of service (uninitialized memory access) or possibly have unspecified other impact via a crafted ELF file.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9653",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "dbus",
						"Namespace": "ubuntu:14.04",
						"Version": "1.6.18-0ubuntu4.3",
						"Vulnerabilities": [
							{
								"Name": "CVE-2015-0245",
								"Namespace": "ubuntu:14.04",
								"Description": "D-Bus 1.4.x through 1.6.x before 1.6.30, 1.8.x before 1.8.16, and 1.9.x before 1.9.10 does not validate the source of ActivationFailure signals, which allows local users to cause a denial of service (activation failure error returned) by leveraging a race condition involving sending an ActivationFailure signal before systemd responds.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-0245",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "zlib",
						"Namespace": "ubuntu:14.04",
						"Version": "1:1.2.8.dfsg-1ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "netbase",
						"Namespace": "ubuntu:14.04",
						"Version": "5.2",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libtasn1-6",
						"Namespace": "ubuntu:14.04",
						"Version": "3.4-3ubuntu0.3",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "sed",
						"Namespace": "ubuntu:14.04",
						"Version": "4.2.2-4ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "e2fsprogs",
						"Namespace": "ubuntu:14.04",
						"Version": "1.42.9-3ubuntu1.3",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "sensible-utils",
						"Namespace": "ubuntu:14.04",
						"Version": "0.0.9",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libusb",
						"Namespace": "ubuntu:14.04",
						"Version": "2:0.1.12-23.3ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "shadow",
						"Namespace": "ubuntu:14.04",
						"Version": "1:4.1.5.1-1ubuntu9.1",
						"Vulnerabilities": [
							{
								"Name": "CVE-2013-4235",
								"Namespace": "ubuntu:14.04",
								"Description": "TOCTOU race conditions by copying and removing directory trees",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2013-4235",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "kmod",
						"Namespace": "ubuntu:14.04",
						"Version": "15-0ubuntu6",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "ucf",
						"Namespace": "ubuntu:14.04",
						"Version": "3.0027+nmu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libpod-latex-perl",
						"Namespace": "ubuntu:14.04",
						"Version": "0.61-1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "ubuntu-keyring",
						"Namespace": "ubuntu:14.04",
						"Version": "2012.05.19",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "perl",
						"Namespace": "ubuntu:14.04",
						"Version": "5.18.2-2ubuntu1",
						"Vulnerabilities": [
							{
								"Name": "CVE-2016-2381",
								"Namespace": "ubuntu:14.04",
								"Description": "environment variable confusion",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-2381",
								"Severity": "Medium",
								"FixedBy": "5.18.2-2ubuntu1.1"
							},
							{
								"Name": "CVE-2013-7422",
								"Namespace": "ubuntu:14.04",
								"Description": "Integer underflow in regcomp.c in Perl before 5.20, as used in Apple OS X before 10.10.5 and other products, allows context-dependent attackers to execute arbitrary code or cause a denial of service (application crash) via a long digit string associated with an invalid backreference within a regular expression.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2013-7422",
								"Severity": "Low",
								"FixedBy": "5.18.2-2ubuntu1.1"
							},
							{
								"Name": "CVE-2014-4330",
								"Namespace": "ubuntu:14.04",
								"Description": "The Dumper method in Data::Dumper before 2.154, as used in Perl 5.20.1 and earlier, allows context-dependent attackers to cause a denial of service (stack consumption and crash) via an Array-Reference with many nested Array-References, which triggers a large number of recursive calls to the DD_dump function.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-4330",
								"Severity": "Low",
								"FixedBy": "5.18.2-2ubuntu1.1"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "lsb",
						"Namespace": "ubuntu:14.04",
						"Version": "4.1+Debian11ubuntu6",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "makedev",
						"Namespace": "ubuntu:14.04",
						"Version": "2.3.1-93ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libterm-ui-perl",
						"Namespace": "ubuntu:14.04",
						"Version": "0.42-1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "isc-dhcp",
						"Namespace": "ubuntu:14.04",
						"Version": "4.2.4-7ubuntu12.3",
						"Vulnerabilities": [
							{
								"Name": "CVE-2015-8605",
								"Namespace": "ubuntu:14.04",
								"Description": "ISC DHCP 4.x before 4.1-ESV-R12-P1 and 4.2.x and 4.3.x before 4.3.3-P1 allows remote attackers to cause a denial of service (application crash) via an invalid length field in a UDP IPv4 packet.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8605",
								"Severity": "Medium",
								"FixedBy": "4.2.4-7ubuntu12.4"
							},
							{
								"Name": "CVE-2016-2774",
								"Namespace": "ubuntu:14.04",
								"Description": "ISC DHCP 4.1.x before 4.1-ESV-R13 and 4.2.x and 4.3.x before 4.3.4 does not restrict the number of concurrent TCP sessions, which allows remote attackers to cause a denial of service (INSIST assertion failure or request-processing outage) by establishing many sessions.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-2774",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 7.1,
											"Vectors": "AV:N/AC:M/Au:N/C:N/I:N"
										}
									}
								}
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "liblog-message-simple-perl",
						"Namespace": "ubuntu:14.04",
						"Version": "0.10-1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "plymouth",
						"Namespace": "ubuntu:14.04",
						"Version": "0.8.8-0ubuntu17.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "audit",
						"Namespace": "ubuntu:14.04",
						"Version": "1:2.3.2-2ubuntu1",
						"Vulnerabilities": [
							{
								"Name": "CVE-2015-5186",
								"Namespace": "ubuntu:14.04",
								"Description": "log terminal emulator escape sequences handling",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-5186",
								"Severity": "Negligible"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libbsd",
						"Namespace": "ubuntu:14.04",
						"Version": "0.6.0-2ubuntu1",
						"Vulnerabilities": [
							{
								"Name": "CVE-2016-2090",
								"Namespace": "ubuntu:14.04",
								"Description": "Heap buffer overflow in fgetwln function of libbsd",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-2090",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "iputils",
						"Namespace": "ubuntu:14.04",
						"Version": "3:20121221-4ubuntu1.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "libgpg-error",
						"Namespace": "ubuntu:14.04",
						"Version": "1.12-0.2ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "pam",
						"Namespace": "ubuntu:14.04",
						"Version": "1.1.8-1ubuntu2",
						"Vulnerabilities": [
							{
								"Name": "CVE-2015-3238",
								"Namespace": "ubuntu:14.04",
								"Description": "The _unix_run_helper_binary function in the pam_unix module in Linux-PAM (aka pam) before 1.2.1, when unable to directly access passwords, allows local users to enumerate usernames or cause a denial of service (hang) via a large password.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-3238",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 5.8,
											"Vectors": "AV:N/AC:M/Au:N/C:P/I:N"
										}
									}
								},
								"FixedBy": "1.1.8-1ubuntu2.1"
							},
							{
								"Name": "CVE-2013-7041",
								"Namespace": "ubuntu:14.04",
								"Description": "The pam_userdb module for Pam uses a case-insensitive method to compare hashed passwords, which makes it easier for attackers to guess the password via a brute force attack.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2013-7041",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 4.3,
											"Vectors": "AV:N/AC:M/Au:N/C:P/I:N"
										}
									}
								},
								"FixedBy": "1.1.8-1ubuntu2.1"
							},
							{
								"Name": "CVE-2014-2583",
								"Namespace": "ubuntu:14.04",
								"Description": "Multiple directory traversal vulnerabilities in pam_timestamp.c in the pam_timestamp module for Linux-PAM (aka pam) 1.1.8 allow local users to create aribitrary files or possibly bypass authentication via a .. (dot dot) in the (1) PAM_RUSER value to the get_ruser function or (2) PAM_TTY value to the check_tty funtion, which is used by the format_timestamp_name function.",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-2583",
								"Severity": "Low",
								"Metadata": {
									"NVD": {
										"CVSSv2": {
											"Score": 5.8,
											"Vectors": "AV:N/AC:M/Au:N/C:P/I:P"
										}
									}
								},
								"FixedBy": "1.1.8-1ubuntu2.1"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "liblockfile",
						"Namespace": "ubuntu:14.04",
						"Version": "1.09-6ubuntu1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "ntp",
						"Namespace": "ubuntu:14.04",
						"Version": "1:4.2.6.p5+dfsg-3ubuntu2.14.04.6",
						"Vulnerabilities": [
							{
								"Name": "CVE-2016-0727",
								"Namespace": "ubuntu:14.04",
								"Description": "NTP statsdir cleanup cronjob insecure",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2016-0727",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-8158",
								"Namespace": "ubuntu:14.04",
								"Description": "Potential Infinite Loop in ntpq",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8158",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-7973",
								"Namespace": "ubuntu:14.04",
								"Description": "Deja Vu: Replay attack on authenticated broadcast mode",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-7973",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-8140",
								"Namespace": "ubuntu:14.04",
								"Description": "ntpq vulnerable to replay attacks",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8140",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-8139",
								"Namespace": "ubuntu:14.04",
								"Description": "Origin Leak: ntpq and ntpdc, disclose origin",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8139",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-7976",
								"Namespace": "ubuntu:14.04",
								"Description": "ntpq saveconfig command allows dangerous characters in filenames",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-7976",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-7979",
								"Namespace": "ubuntu:14.04",
								"Description": "Off-path Denial of Service (DoS) attack on authenticated broadcast mode",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-7979",
								"Severity": "Low"
							},
							{
								"Name": "CVE-2015-8138",
								"Namespace": "ubuntu:14.04",
								"Description": "ntp: missing check for zero originate timestamp",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-8138",
								"Severity": "Medium"
							},
							{
								"Name": "CVE-2015-7977",
								"Namespace": "ubuntu:14.04",
								"Description": "reslist NULL pointer dereference",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-7977",
								"Severity": "Medium"
							},
							{
								"Name": "CVE-2015-7978",
								"Namespace": "ubuntu:14.04",
								"Description": "Stack exhaustion in recursive traversal of restriction list",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-7978",
								"Severity": "Medium"
							},
							{
								"Name": "CVE-2015-7974",
								"Namespace": "ubuntu:14.04",
								"Description": "NTP 4.x before 4.2.8p6 and 4.3.x before 4.3.90 do not verify peer associations of symmetric keys when authenticating packets, which might allow remote attackers to conduct impersonation attacks via an arbitrary trusted key, aka a \"skeleton key.\"",
								"Link": "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2015-7974",
								"Severity": "Low"
							}
						],
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "newt",
						"Namespace": "ubuntu:14.04",
						"Version": "0.52.15-2ubuntu5",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					},
					{
						"Name": "sqlite3",
						"Namespace": "ubuntu:14.04",
						"Version": "3.8.2-1ubuntu2.1",
						"AddedBy": "sha256:d89e1bee20d9cb344674e213b581f14fbd8e70274ecf9d10c514bab78a307845"
					}
				]
			}
		},
		{
			"Layer": {
				"Name": "sha256:9e0bc8a71bde464f710bc2b593a1fc21521517671e918687892303151331fa56",
				"Namespace": "ubuntu:14.04",
				"ParentName": "sha256:27aa681c95e5165caf287dcfe896532df4ae8b10e099500f2f8f71acf4002a89",
				"IndexedByVersion": 2
			}
		},
		{
			"Layer": {
				"Name": "sha256:27aa681c95e5165caf287dcfe896532df4ae8b10e099500f2f8f71acf4002a89",
				"Namespace": "ubuntu:14.04",
				"ParentName": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4",
				"IndexedByVersion": 2
			}
		},
		{
			"Layer": {
				"Name": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4",
				"IndexedByVersion": 2
			}
		}
	]
}
`