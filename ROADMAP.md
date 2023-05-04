| Branch or tag | Feature                                               | Status/Comments                                                                                                              |
|---------------|-------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------|
| 0.100         | initial groundwork                                    | *completed*                                                                                                                  |
| 0.200         | `ca create`, `ca verify`                              | *completed*                                                                                                                  |
| 0.300         | `ca edit`, `ca del`                                   | *completed* --> `ca edit` is removed. There is no such thing as an edited cert. You delete it  and create a new one, period. |
| 0.400         | environment config management                         | *underway*                                                                                                                   |
| 0.500         | `cert create`, `cert verify`, `cert edit`, `cert del` | will reuse most of the `ca` code                                                                                             |
| 0.600         | java ssl certificates (JKS)                           | needs a way to encrypt-decrypt the password in the JSON file                                                                 |
| 0.700         | create eco-system (subdirs) files                     | directories for such as private keys, crl, index, etc                                                                        |
| 0.800         | portability                                           | replace all paths with '/' with filepath.Join()                                                                              |
| 0.900         | major package refactoring                             | might or might not... this will most likely be the last branch before release                                                |
| 1.000         | release                                               |
| 1.000+        | `ca install`, `ca revoke`                             |



