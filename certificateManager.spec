%define debug_package   %{nil}
%define _build_id_links none
%define _name   certificateManager
%define _prefix /opt
%define _version 0.200
%define _rel 0
%define _arch x86_64

Name:       certificateManager
Version:    %{_version}
Release:    %{_rel}
Summary:    certificateManager

Group:      SSL
License:    GPL2.0
URL:        https://github.com/jeanfrancoisgratton/certificateManager

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: gcc
#Requires: sudo
#Obsoletes: vmman1 > 1.140

%description
RootCA and server SSL certificate manager

%prep
#%setup -q
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_name} .
strip %{_sourcedir}/%{_name}

%clean
rm -rf $RPM_BUILD_ROOT

%pre
exit 0

%install
#%{__mkdir_p} "$RPM_BUILD_ROOT%{_prefix}/bin"
#install -Dpm 0755 %{buildroot}/%{_name} "$RPM_BUILD_ROOT%{_prefix}/bin/"
install -Dpm 0755 %{_sourcedir}/%{name} %{buildroot}%{_bindir}/%{name}

%post

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{name}


%changelog
* Sun Apr 16 2023 builder <builder@famillegratton.net> 0.101-0
- Version bump and fixes (jean-francois@famillegratton.net)

* Sun Apr 16 2023 builder <builder@famillegratton.net> 0.100-0
- Initial dry-run on packaging

