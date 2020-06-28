**_Note:_** Because I'm running my own servers for several years, main development is done at at https://git.ypbind.de/cgit/check_nebula_certificate_expiration/

----

# Preface
[Nebula VPN](https://github.com/slackhq/nebula) is a scalable and portable VPN mesh using certificates to encrypt and authenticate
clients. Because Nebula VPN uses it's own certificate format and does not uses standard x509 certificates normal
checks for the expiration of a certificate are not possible.

This check allows for checking the expiration of a Nebula VPN certificate and can be integrated
in a Nagios based monitoring solution like [Icinga2](https://icinga.com/products/)

# Build requirements
This check is implemented in Go so the Go compiler is required.
To build this programm the nebula (https://github.com/slackhq/nebula/) package is required.
Formatting of the duration is done by a slightliy improved version of the [durafmt](https://github.com/hako/durafmt) included in this source code.

# Command line parameters
The path to the Nebula VPN certificate (`--certificate`) is mandatory. Warn and critical thresholds can be specified using a format suitable for [time.Duration.ParseDuration](https://golang.org/pkg/time/#ParseDuration).

| *Paraemter* | *Description* | *Default* | *Comment* |
|:------------|:--------------|:---------:|:----------|
| `--help` | Show help text | - | - |
| `--certificate=<file>` | Certificate file to check | - | **Mandatory** |
| `--critical=<time>` | Report critical status if certificate expires in <time> or less | 168h | - |
| `--version` | Show version information | - | - |
| `--warning=<time>` | Report warning status if certificate expires in <time> or less | 336h | - |

# License
## check_nebula_certificate_expiration
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

## durafmt (https://github.com/hako/durafmt)
The MIT License (MIT)

Copyright (c) 2016 Wesley Hill

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

## nebula (https://github.com/slackhq/nebula/)
MIT License

Copyright (c) 2018-2019 Slack Technologies, Inc.

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

