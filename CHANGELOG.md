## rmapi master

## rmapi 0.0.24 (December 30, 2022)

- Make sure Windows binaries have an exe extension (@grahamgreen)

- Use go 1.19 (@grahamgreen)

## rmapi 0.0.23 (December 01, 2022)

- sync 1.5: use parallel requests when getting remote tree (@ddvk)

- fix syncomplete in mput (@ddvk)

- update auth url (@ddvk)

- Add support to nuke 1.5 (@ddvk)


## rmapi 0.0.22 (October 05, 2022)

- Update 1.5 sync URLs, implement v2 API (@ddvk)

## rmapi 0.0.21 (September 04, 2022)

- Fix time conversion (@ddvk)

- Update install instructions (@ddvk)

- Tweak export brush size (@pacien)

- Update x/sys dep to fix builds with newer Go versions (@andrew-d)

- Use Go 1.18

## rmapi 0.0.20 (July 6, 2022)
- Fix timestamps for sync 1.5 (@ddvk)

- Output warnings for to stderr for sync 1.5 (@ddvk)

- Some hack with pages (@abmantis)

## rmapi 0.0.19 (January 29, 2022)

- Fix issue with rate limiting and mput (@ddvk)

- Fix type in help test (@myersjustinc)

## rmapi 0.0.18 (November 15, 2021)

* More fixes for new sync protocol (@ddvk)

## rmapi 0.0.17 (November 02, 2021)

* Some fixes for new sync protocol (@ddvk)

## rmapi 0.0.16 (October 24, 2021)

* Use Go 1.16

* Fixed OTP url in CLI and docs

* Add experimental support for new sync protocol

## rmapi 0.0.15 (May 19, 2021)

* Fix authentication URL

* stats outputs JSON now

## rmapi 0.0.14 (May 03, 2021)

* Don't dump request/responses when tracing disabled
  to fix OOM on low end devices (@ddvk)

* Fix small typo (Jasdev Singh)

* Show entry's type in find command output (@ddvk)

* Remove device token if unable to authenticate (@ddvk)

* Upload .rm page files (@ddvk)

* Fix panic in mv command (Casey Marshall)

* Update install instructions with config folder command (@Caleb)

* Update doc to correct OTC URL (Hwee-Boon Yar)

* Fix BrushSize parsing (@jgoel)

## rmapi 0.0.13 (December 08, 2020)

* Copy table contents when creating annotations (@ddvk)

* Incremental sync (@ddvk)

* Fix auth retries (@ddvk)

* Use nested config dir under XDG (@jebw)

* Bump go verison to 1.15 (@grahamgreen)

## rmapi 0.0.12 (June 18, 2020)

* Use XDG paths conf config file (@ddvk)

* Fix issue where documents are downloaded again
  when the device reboots (@ddvk)

* Fix annnoation issue where Acrobat Reader wouldn't display
  annotations correctly (@ddvk)

## rmapi 0.0.11 (April 28, 2020)

* Add env variables to override cloud API hosts (@ddvk)

* Upload downloaded zip files (@ddvk)

* Bug fix: use UTC time when setting document's upload time (@ddvk)

* Add support to optinally create thumbnails for uploaded documents (@ddvk)

* Update CI scripts to use Go 1.13 (@ddvk)

## rmapi 0.0.10 (April 17, 2020)

* Multiple annotation fixes (@ddvk)

* Add support to create thumbnails in large PDF docs (@ddvk)

* Use community license instead a UniPDF fork (@ddvk)

* Fix put bug to allow directories and files with the same name (@GjjvdBurg)

## rmapi 0.0.9 (February 01, 2020)

* Change license to AGPL

* Initial support for PDF annotations with UniPDF

## rmapi 0.0.8 (January 06, 2020)

* Add support for v5 annotations

## rmapi 0.0.7 (November 11, 2019)

* Rename api.v2 to cloud

## rmapi 0.0.6 (September 08, 2019)

* Migrate to go11 modules

* Add api.v2

## rmapi 0.0.5 (August 11, 2019)

* Fix issue to make document uploads work with reMarkable 1.7.x firmware upgrade

* Increased http timeout to 5 minutes to enable upload of larger files

* Add user-agent header to be a good reMarkable citezen

* Ls may take a directory as an argument

* Ignore hidden files/directories by default

* Initial support for annotations

* Fix panic when autocompleting the "ls" command

* Add find command

## rmapi 0.0.4 (October 1, 2018)

* Windows fixes

* Add autocompleter for local files that is used by "put"

* Fix mv command

* Put may take a second argument as destination

* Autocomplete for "put" command only shows ".pdf" files

* Add support to upload epub files

* rm supports multiple files

* Return exit code for non-interactivly commands

* Vendorize fuse dependencies

* Use new auth endpoints

## rmapi 0.0.3 (February 25, 2018)

* Update doc

* Fix file upload

   *Javier Uruen Val*

## rmapi 0.0.2 (February 25, 2018)

*  Fix directory creation (fixes #6)

*  Add stat command to show entry's metadata

   *Javier Uruen Val*

## rmapi 0.0.1 (February 24, 2018)

*   Initial release with support for most of the API and autocompletion.

    *Javier Uruen Val*
