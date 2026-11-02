# Protocol

The console asks Sony for a plain-HTTP XML list of the newest system software.

## Hosts and path

`http://<d|f|h><region>01.ps4.update.playstation.net/update/ps4/list/<region>/ps4-updatelist.xml`

Regions: `au br cn eu hk jp kr mx ru sa tw uk us`. Examples: `dus01`, `fus01`, `duk01`. The `d`, `f` and
`h` prefixes are separate Sony hosts for the same content.

## Observed response (firmware 14.00, region us, 2026-11-02)

Lines end in CRLF.

```xml
<?xml version="1.0" ?>
<update_data_list>
	<region id="us">
		<force_update>
			<system level0_system_ex_version="0" level0_system_version="13.520.000" level1_system_ex_version="0" level1_system_version="13.520.000"/>
		</force_update>
		<system_pup ex_version="0" label="14.00" sdk_version="14.008.001" version="14.008.000">
			<update_data update_type="full">
				<image size="503331328">http://dus01.ps4.update.playstation.net/update/ps4/image/2026_0831/sys_8b844e5e5bb63ed19ef64b8971b400b0/PS4UPDATE.PUP?dest=us</image>
			</update_data>
		</system_pup>
		<recovery_pup type="default">
			<system_pup ex_version="0" label="14.00" sdk_version="14.008.001" version="14.008.000"/>
			<image size="1083638784">http://dus01.ps4.update.playstation.net/update/ps4/image/2026_0831/rec_f6a3206393b32855b8bf533cc5a4d29a/PS4UPDATE.PUP?dest=us</image>
		</recovery_pup>
	</region>
</update_data_list>
```

`force_update` is the minimum version a console must install; `system_pup` is the newest version, with
`update_data` its full update image; `recovery_pup` is the reinstall image. Image URLs carry `?dest=<region>`.

## What we change

- The region in `<region id>` echoes the request path.
- Version, label and SDK version come from `PS4_UPDATE_VERSION`, `PS4_UPDATE_LABEL`, `PS4_UPDATE_SDK_VERSION`.
  The default `00.000.000` is older than any console's firmware, so none sees an update.
- With `PS4_UPDATE_IMAGE_URL` empty, `update_data` and `recovery_pup` are omitted and the force levels are
  `00.000.000`. With it set, both images use that URL plus `?dest=<region>` and the force levels equal the version.
  Image `size` is always `0`.
