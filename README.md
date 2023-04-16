<H1>certificateManager</H1>
___
Tool to manage your own certificates<br><br>

With this tool you will be able to:
- manage both root CA and server certificates
  - Create new root CA
  - Remove (revoke) old server certs
  - Creste new server certs against the rootCA

- create the configuration files needed for the above tasks

<H2>Installing from source</H2>
<H3>Pre-requisites:</H3>
You need the following packages and tools:
- gcc
- make
- go : check the `go.version` file to see which version to install; as of writing this, v1.20.3 was the current version.

<H3>Building:</H3>
A helper script is available in the `src/` directory.<br>
You actually need to go into that directory to build the binary, and then just run: `./build.sh`. The binary will be copied to `/opt/bin` by default, but this path can be changed by adding a parameter to the script (ex: `./build.sh /usr/local/bin`).

Optional: you might wish to strip the binary from debugging code once installed; you do this like this: `strip /opt/bin/certificateManager` (assuming of course that the binary is installed in `/opt/bin`)


<H2>Installing from a repository</H2>
This software is bundled as a RedHat/OpenSUSE unsigned RPM, as a Debian/Ubuntu .DEB, and as an Alpine APK.

For all of these formats, you will need to install my Root CA certificate, located under `$ROOTDIR/rootCert/rootCAfamillegratton.crt`

<H3>RedHat/CentOS/Fedora/RockyLinux/OpenSUSE</H3>
Copy the following block as `famillegratton.repo`; the file should be in `/etc/yum.repos.d/` for most RPM-based distro, and in `/etc/zypp/repos.d` for OpenSUSE.

```
[famillegratton]
enabled=1
autorefresh=1
baseurl=https://nexus.famillegratton.net:1808/repository/dnfLocal
```

<H3>Debian/Ubuntu</H3>
Create the following file under `/etc/apt/sources.list.d/famillegratton.list` :

```
deb [arch=amd64] https://nexus.famillegratton.net:1808/repository/aptLocal nexus main
```

Refresh your repo lists (`dnf clean all && dnf makecache`, `apt-get update`, `zypper ref`), and then use your usual tool to download and install

<H2>How to use the software</H2>
==> First, you will need to generate a skeleton of configuration file:
`certificateManager config CAtemplate` :
This will create a skeleton file to generate custom rootCAs under `$HOME/.config/certificateManager/rootCA-default.json`
Edit this file as needed.

==> Then, if you need server certificates signed against that rootCA, you will need to generate its template, too:
`certificateManager config Certtemplate` . The config file will be named `$HOME/.config/certificateManager/serverCert-default.json`

This software allows you to use different config files at runtime with the `-e` flag. For instance, if I wanted to run the tool using the config for a server named mediaserver1, you would something like this:
`certificateManager -e mediaserver1 OTHER_COMMANDS`. This would fetch all of its required info from `$HOME/.config/certificateManager/mediaserver1.json`

<H3>Create a custom root CA</H3>
Assuming that you have a config file named `rootca.json` :<br>
`certificateManager -e rootca ca create`<br>
Please have a look at the different flags you could use: `certificateManager ca -h` or `certificateManager ca create -h`

<H3>Create a server certificate signed against your own root CA</H3>
Assuming you have a config file named `server.json` :<br>
`certificateManager -e server.json cert create`. Again, `certificateManager cert -h` is your friend.


As a matter of fact, this software uses GO's COBRA-CLI framework, a very arg parser commonly used by Docker, Kubernetes, Terraform. You can get help from many, if not all commands.
