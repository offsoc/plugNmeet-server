# Changelog

## [1.10.0](https://github.com/offsoc/plugNmeet-server/compare/v1.9.1...v1.10.0) (2025-08-26)


### Features

* dependencies injection using wire ([0d6ddef](https://github.com/offsoc/plugNmeet-server/commit/0d6ddefc6b718d7c546ab0b582c21ed46062cf02))
* dependencies injection using wire ([a9e78ff](https://github.com/offsoc/plugNmeet-server/commit/a9e78ff03942386ce3d6177f8a7e8aabbd2f0a17))
* extra_data option for user ([70d032d](https://github.com/offsoc/plugNmeet-server/commit/70d032dedfebdeec161bfbaee8ed7a5c2b43544b))
* option to auto generate user_id ([89f5fb7](https://github.com/offsoc/plugNmeet-server/commit/89f5fb7758d589d123b03e6e8ae8c4bcaf0ce8a6))
* option to encrypt chat messages ([8f877e9](https://github.com/offsoc/plugNmeet-server/commit/8f877e98fc8eaacdadfebf4ea29f2c21453a74fd))
* option to keep deleted recordings in backup ([7d1ee5d](https://github.com/offsoc/plugNmeet-server/commit/7d1ee5db73d3689d7137233ba0dbe48294eb82a6))
* option to keep deleted recordings in backup for certain period ([cf64f3c](https://github.com/offsoc/plugNmeet-server/commit/cf64f3ce982fe97096796b4f7b50f755a7ec6a18))
* option to set `max_size_whiteboard_file` ([1ae3df8](https://github.com/offsoc/plugNmeet-server/commit/1ae3df857bfecdd6a29a10b429be66b583008d3e))
* option to set log level ([816d450](https://github.com/offsoc/plugNmeet-server/commit/816d450364a611a442d0c64af9e957f675b694d7))
* replaced websocket with nats ([#550](https://github.com/offsoc/plugNmeet-server/issues/550)) ([393e124](https://github.com/offsoc/plugNmeet-server/commit/393e124b081cbed3d8b564c38b34fa9b9d4e2e36))


### Bug Fixes

* `ex_user_id` was changing because of pointer ([73a5d98](https://github.com/offsoc/plugNmeet-server/commit/73a5d98b2da8ea2e8b792cc4aabfef5e6e3977ee))
* added option `-trimpath` ([6b355d9](https://github.com/offsoc/plugNmeet-server/commit/6b355d9a89c668e400fefbcfea4dcfae54e14bee))
* added option to insert a self provided e2ee key ([386cbdd](https://github.com/offsoc/plugNmeet-server/commit/386cbdd1ce79e32ce206f18164fc7ed84abffc64))
* api to upload base64 encoded files ([f7b9923](https://github.com/offsoc/plugNmeet-server/commit/f7b992311ab4ef0a0214030444bb8fd1b420ca43))
* bug + clean up ([45bca50](https://github.com/offsoc/plugNmeet-server/commit/45bca50d19638a3f7e273d2a38f6d7fdd05f8bbc))
* **bug:** Ingress was not working in the new Nats solution. Ref: https://github.com/mynaparrot/plugNmeet-server/discussions/611 ([2fc55a4](https://github.com/offsoc/plugNmeet-server/commit/2fc55a4e5c02787962a4e77804cfb83b38c7d975))
* **bug:** ingress was removing because of invalid inactivity checking ([c08ec4a](https://github.com/offsoc/plugNmeet-server/commit/c08ec4ab3f060183c328d0d40a692ab8ab455ce6))
* **bug:** user's online status wasn't updated properly in the latest version of nats-server ([f9d9793](https://github.com/offsoc/plugNmeet-server/commit/f9d97932e902832455f7f4ce6e13bd8f702679ab))
* bump nodejs ([2779eab](https://github.com/offsoc/plugNmeet-server/commit/2779eab3538a008d693e9a577b07aa78c5640a0e))
* bump proto ([acb2bbf](https://github.com/offsoc/plugNmeet-server/commit/acb2bbf9a0d80041483313c3404189a813795f1e))
* bump proto ([21b0c2d](https://github.com/offsoc/plugNmeet-server/commit/21b0c2de01356ad588d53244c51e17d788d98214))
* bump proto ([9fcde46](https://github.com/offsoc/plugNmeet-server/commit/9fcde46144b7a0511dc77d218276b5f32f99e97b))
* **bump:** bump version ([de8f4dd](https://github.com/offsoc/plugNmeet-server/commit/de8f4dd1f0dd5086684738c7d762360cfccd28d8))
* **ci:** added beta release ([e7497e2](https://github.com/offsoc/plugNmeet-server/commit/e7497e2a84f3768007ecb1a787feea0ee87b883c))
* **ci:** always-update ([1134c50](https://github.com/offsoc/plugNmeet-server/commit/1134c50f91c84c7fc9eabf5e82f9b50ffdc01472))
* **ci:** bump beta version ([b79a0ea](https://github.com/offsoc/plugNmeet-server/commit/b79a0eadd93a0e06070f12dc087c4f18ec49247a))
* **ci:** bump beta version ([24c9082](https://github.com/offsoc/plugNmeet-server/commit/24c90821b4e9309ec6666ddcc19caf0f6069168d))
* **ci:** removed rebase again ([db27ad5](https://github.com/offsoc/plugNmeet-server/commit/db27ad54d16a4b39c9b5347ee2f06a8164a95266))
* **ci:** To rebase again ([e1681fc](https://github.com/offsoc/plugNmeet-server/commit/e1681fc79aede64a5a9cc297d7a4505461767eb8))
* **ci:** To rebase again ([161b04e](https://github.com/offsoc/plugNmeet-server/commit/161b04ed775d25a873617c229a0a8bfb23989672))
* **ci:** To rebase again ([4a34670](https://github.com/offsoc/plugNmeet-server/commit/4a3467056d97d57245219b0f7cb6d406c3a1fd92))
* clean up ([6e12ff1](https://github.com/offsoc/plugNmeet-server/commit/6e12ff16e7e5dea4e2c41c64dee5e1e8cdb13053))
* cleaned up logic ([4aa5b49](https://github.com/offsoc/plugNmeet-server/commit/4aa5b4930286af2d5379cafc3b914bcf36cde083))
* **cleanup:** remove unnecessary goroutine ([b5f6d0a](https://github.com/offsoc/plugNmeet-server/commit/b5f6d0a82a232ad3d97ef82ffe4bfe47a86407ac))
* delay of user disconnect event ([ba30ad4](https://github.com/offsoc/plugNmeet-server/commit/ba30ad4579f4fbed5969ab4185e673a6e1887551))
* deps update ([264d0ee](https://github.com/offsoc/plugNmeet-server/commit/264d0eecf750ece8ebc77cd8361997003e4df019))
* deps update ([93fff85](https://github.com/offsoc/plugNmeet-server/commit/93fff85fe110fc79da5a602c51d2f91c8031e321))
* deps update ([0d58c13](https://github.com/offsoc/plugNmeet-server/commit/0d58c1323a618e4fe6c2bdcb2fd853409851e3a5))
* **deps:** dependencies update ([ef83e0d](https://github.com/offsoc/plugNmeet-server/commit/ef83e0daaecb76dc884f53c4c5187d65619d3968))
* **deps:** deps update ([90d40b3](https://github.com/offsoc/plugNmeet-server/commit/90d40b3a6faa84da20889687d286fa62acb13b0e))
* **deps:** deps update ([a03f37e](https://github.com/offsoc/plugNmeet-server/commit/a03f37e3cc9efa3baf7c7866157460c27350c1b7))
* **deps:** deps update ([3d8d4b9](https://github.com/offsoc/plugNmeet-server/commit/3d8d4b933cdb6fd8a992ddc672d22c83795a2be6))
* **deps:** protocol update ([a321514](https://github.com/offsoc/plugNmeet-server/commit/a32151457d29ff33ca8fd5be7a27f9f12a731c50))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 0b47c3e ([1ba7085](https://github.com/offsoc/plugNmeet-server/commit/1ba70850485d3b8f9d5ef232b7243d9f7eb5f5ea))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 225c77f ([8e5d44e](https://github.com/offsoc/plugNmeet-server/commit/8e5d44e14cee51ea25b5e599973a746d4ff0a707))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 225c77f ([d65403f](https://github.com/offsoc/plugNmeet-server/commit/d65403f1c67982b3aebfd1164d71ca33c0d52223))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 4d9a8f6 ([ca47af5](https://github.com/offsoc/plugNmeet-server/commit/ca47af52c74c9644c8d3ee5ee00a9e256f855dd6))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 4d9a8f6 ([d9ec4b4](https://github.com/offsoc/plugNmeet-server/commit/d9ec4b43c9abebe5ee05a59ffc9e01b417090691))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 5a79e5a ([0a642d1](https://github.com/offsoc/plugNmeet-server/commit/0a642d121a2e7d3112202fdf7d482d0dbfe3063b))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 5a79e5a ([b009cba](https://github.com/offsoc/plugNmeet-server/commit/b009cba839ef641cd5a987b42b986589636c055d))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 5f7afa8 ([ed532cd](https://github.com/offsoc/plugNmeet-server/commit/ed532cdd3a9c754d742fc70c641374934f695b5f))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 5f7afa8 ([0e915fc](https://github.com/offsoc/plugNmeet-server/commit/0e915fc686d726bcc7061195787eb5d6ab5a25d4))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 73eee1f ([8e9ca0a](https://github.com/offsoc/plugNmeet-server/commit/8e9ca0aa564b4b34ac9d478b5253f83667befb73))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 73eee1f ([05b8d9f](https://github.com/offsoc/plugNmeet-server/commit/05b8d9f50a96fbfb6681abed793b0072e91ce295))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 76539e3 ([2bebc97](https://github.com/offsoc/plugNmeet-server/commit/2bebc9752c5c841daa6d718cba2fe584e2ccc62b))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 76539e3 ([954552b](https://github.com/offsoc/plugNmeet-server/commit/954552bc97da21a00da9ab117f2c1d8b1dc9fa06))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to b2a5cc6 ([ce7e129](https://github.com/offsoc/plugNmeet-server/commit/ce7e12944665765bc7918d6fe1d5425bd413748a))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to b2a5cc6 ([3e40713](https://github.com/offsoc/plugNmeet-server/commit/3e4071398f9c55b6e305a242f789f0d53126e3e4))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to eed083b ([fd61a85](https://github.com/offsoc/plugNmeet-server/commit/fd61a859727e2261d235421efda16204e6e57548))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to eed083b ([f5d0681](https://github.com/offsoc/plugNmeet-server/commit/f5d0681369e71d7e92c8ff6cf1291ab7241eeaaf))
* **deps:** update module github.com/ansrivas/fiberprometheus/v2 to v2.10.0 ([7364863](https://github.com/offsoc/plugNmeet-server/commit/7364863bad7628530fce3f17a8a40fca038a6ce0))
* **deps:** update module github.com/ansrivas/fiberprometheus/v2 to v2.10.0 ([cfbb9f3](https://github.com/offsoc/plugNmeet-server/commit/cfbb9f336e38382a4802a700b194f0cc92b5e81c))
* **deps:** update module github.com/ansrivas/fiberprometheus/v2 to v2.12.0 ([b5c2b26](https://github.com/offsoc/plugNmeet-server/commit/b5c2b26cc02fc0be77bbdf660350fccb9a0ab04e))
* **deps:** update module github.com/ansrivas/fiberprometheus/v2 to v2.12.0 ([3bda07e](https://github.com/offsoc/plugNmeet-server/commit/3bda07e3d1b7abd4f863547e21aca9a2245d03c6))
* **deps:** update module github.com/ansrivas/fiberprometheus/v2 to v2.7.0 ([7377bd5](https://github.com/offsoc/plugNmeet-server/commit/7377bd5466391e0c30e7c314fe712b4841ad1d0b))
* **deps:** update module github.com/ansrivas/fiberprometheus/v2 to v2.7.0 ([4ea74d1](https://github.com/offsoc/plugNmeet-server/commit/4ea74d15e944b8ae2bf7a9d58bdb09f0d8aa9559))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.6.2 ([6213d98](https://github.com/offsoc/plugNmeet-server/commit/6213d98e776c1eb1ecd285831fb261ef570aaa26))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.6.2 ([c3bd6a5](https://github.com/offsoc/plugNmeet-server/commit/c3bd6a53933418e201d1b26dcff24033906b15e1))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.6.3 ([3ba52f5](https://github.com/offsoc/plugNmeet-server/commit/3ba52f5ffc7214b881f586c5f4b08ebc98c4ef1f))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.6.3 ([c061992](https://github.com/offsoc/plugNmeet-server/commit/c0619925465bed6d7350c3457f2aaba962674355))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.7.2 ([e0f2dbe](https://github.com/offsoc/plugNmeet-server/commit/e0f2dbe1e26ffb55fafc2fc94adfb732606b38a6))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.7.2 ([118e56a](https://github.com/offsoc/plugNmeet-server/commit/118e56a40302ae98795984ce29f51d641b2f6486))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.7.3 ([e740d5a](https://github.com/offsoc/plugNmeet-server/commit/e740d5a5a04091cf327f8d4fd09f04e0b2b74f59))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.7.3 ([0c51ba8](https://github.com/offsoc/plugNmeet-server/commit/0c51ba89ad3037b71bf5a3f7ae83272811c0a78e))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.9.2 ([6052e9e](https://github.com/offsoc/plugNmeet-server/commit/6052e9e64488d51e0f74c36b81ecd6cc8e4195c5))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.9.2 ([e924e7b](https://github.com/offsoc/plugNmeet-server/commit/e924e7b1f0bb481b367b62f794438b1e3b5098c1))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.4 ([c9b858b](https://github.com/offsoc/plugNmeet-server/commit/c9b858b77e736cc7ae28ec13158fbc894c5ae90a))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.4 ([3143831](https://github.com/offsoc/plugNmeet-server/commit/3143831cc811bfcb58b9b02a30581d7aee16f2aa))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.5 ([e52f6b6](https://github.com/offsoc/plugNmeet-server/commit/e52f6b6345670ef2c62ef7340c07ddfbc85ab536))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.5 ([611534e](https://github.com/offsoc/plugNmeet-server/commit/611534e766f943935c9238ea13bb8e6bc6229659))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.6 ([7100da0](https://github.com/offsoc/plugNmeet-server/commit/7100da09e9b3ff590193e07bbb112ffddb786c92))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.6 ([fc6df57](https://github.com/offsoc/plugNmeet-server/commit/fc6df57253b8ff035cc87e3d76d651e2391a9419))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.7 ([6c95ae5](https://github.com/offsoc/plugNmeet-server/commit/6c95ae55d91ae3bb1f294e2a3847454ec44590cc))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.7 ([56f5d3b](https://github.com/offsoc/plugNmeet-server/commit/56f5d3becfacba5c535cf45615c2fdeaf898274d))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.8 ([5a09c18](https://github.com/offsoc/plugNmeet-server/commit/5a09c18f581fb2e56468fd97f43d62c3c690115b))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.8 ([703623d](https://github.com/offsoc/plugNmeet-server/commit/703623d55bbcd8a2b9c613dbe07446670e3f2ccb))
* **deps:** update module github.com/go-jose/go-jose/v4 to v4.0.5 [security] ([42be81b](https://github.com/offsoc/plugNmeet-server/commit/42be81ba43c996af7d0ee9bcb4cfe79b1c11fd21))
* **deps:** update module github.com/go-jose/go-jose/v4 to v4.0.5 [security] ([084bc36](https://github.com/offsoc/plugNmeet-server/commit/084bc367dbf9ac2fa021ccf0fb42caa651cbd01d))
* **deps:** update module github.com/goccy/go-json to v0.10.3 ([f4db427](https://github.com/offsoc/plugNmeet-server/commit/f4db427a87c9d8e72150c7d39771cd4a4c5fdf45))
* **deps:** update module github.com/goccy/go-json to v0.10.3 ([ee184ef](https://github.com/offsoc/plugNmeet-server/commit/ee184effa44cfaba81f2604054c63515a6a8a1ca))
* **deps:** update module github.com/goccy/go-json to v0.10.5 ([c754be0](https://github.com/offsoc/plugNmeet-server/commit/c754be0a983e025d03887f5851773d5f02a86600))
* **deps:** update module github.com/goccy/go-json to v0.10.5 ([93ef56e](https://github.com/offsoc/plugNmeet-server/commit/93ef56eeeb28292a1d0409684bc71a624bd414aa))
* **deps:** update module github.com/gofiber/contrib/socketio to v1.0.1 ([f699a76](https://github.com/offsoc/plugNmeet-server/commit/f699a763c0f3f190aa949683d25d5f5eb21c7402))
* **deps:** update module github.com/gofiber/contrib/socketio to v1.0.1 ([2c4c856](https://github.com/offsoc/plugNmeet-server/commit/2c4c85641852231f7ac4d54e4c36b3c0714f6409))
* **deps:** update module github.com/gofiber/contrib/socketio to v1.1.1 ([ce94c9c](https://github.com/offsoc/plugNmeet-server/commit/ce94c9c79a70ee0e3ca09f435cbb89a683c98a37))
* **deps:** update module github.com/gofiber/contrib/socketio to v1.1.1 ([bc23ce7](https://github.com/offsoc/plugNmeet-server/commit/bc23ce7d88bd249aa67ab5991f1d7578d91937a3))
* **deps:** update module github.com/gofiber/contrib/socketio to v1.1.2 ([7a1b99f](https://github.com/offsoc/plugNmeet-server/commit/7a1b99fec9ba9be1cdde86e73bac62bfde7302eb))
* **deps:** update module github.com/gofiber/contrib/socketio to v1.1.2 ([3587f83](https://github.com/offsoc/plugNmeet-server/commit/3587f83d511e5fd24531c6430ecc7d6b06eb28b2))
* **deps:** update module github.com/gofiber/contrib/websocket to v1.3.1 ([36dbce8](https://github.com/offsoc/plugNmeet-server/commit/36dbce8b8b0c6499dde04f407ade54e785a84b28))
* **deps:** update module github.com/gofiber/contrib/websocket to v1.3.1 ([c7b0fa5](https://github.com/offsoc/plugNmeet-server/commit/c7b0fa5b2e287caf756bf9550fdad0a47297417f))
* **deps:** update module github.com/gofiber/contrib/websocket to v1.3.2 ([eb0bda1](https://github.com/offsoc/plugNmeet-server/commit/eb0bda11c80753e680aa8e34dbca62ff5fd5c084))
* **deps:** update module github.com/gofiber/contrib/websocket to v1.3.2 ([d4c8623](https://github.com/offsoc/plugNmeet-server/commit/d4c8623ef0d22b434d96420742c53b33371240a4))
* **deps:** update module github.com/gofiber/fiber/v2 to v2.52.5 ([070c311](https://github.com/offsoc/plugNmeet-server/commit/070c311dd617e5d75c86de608e041867851ef850))
* **deps:** update module github.com/gofiber/fiber/v2 to v2.52.5 ([b5bd2ed](https://github.com/offsoc/plugNmeet-server/commit/b5bd2ed84d103128f12f434e8abc1c9954ba5768))
* **deps:** update module github.com/gofiber/template/html/v2 to v2.1.2 ([0b25c31](https://github.com/offsoc/plugNmeet-server/commit/0b25c31aaa6d88d246296cfd2e5a5ed3f0f58aa2))
* **deps:** update module github.com/gofiber/template/html/v2 to v2.1.2 ([f1ad262](https://github.com/offsoc/plugNmeet-server/commit/f1ad262b0fb13a2a861eef8c090b06206755280f))
* **deps:** update module github.com/livekit/protocol to v1.27.1 ([85682c7](https://github.com/offsoc/plugNmeet-server/commit/85682c75b0b7d453472838525b90865b8f97efdb))
* **deps:** update module github.com/livekit/protocol to v1.27.1 ([373f1ba](https://github.com/offsoc/plugNmeet-server/commit/373f1ba67264263816b1fb0a912f935d44781e5e))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.1.2 ([f30ff87](https://github.com/offsoc/plugNmeet-server/commit/f30ff870ca5e4a7b361aa3f9310c944bea849a8f))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.1.2 ([b6c5c65](https://github.com/offsoc/plugNmeet-server/commit/b6c5c653e1a01c819380dbcc6afe52a286f07490))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.1.3 ([41139ec](https://github.com/offsoc/plugNmeet-server/commit/41139ecc7f72e6ccdfa88495dd4224bfb67f9674))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.1.3 ([6301045](https://github.com/offsoc/plugNmeet-server/commit/63010451d76f1ba879567d620d3f85fb6dd1bcc0))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.2.0 ([e6284e5](https://github.com/offsoc/plugNmeet-server/commit/e6284e59edd594411c8e2d6a0fae8e8d3d6f2f5d))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.2.0 ([bbe3618](https://github.com/offsoc/plugNmeet-server/commit/bbe3618170edb48f77140115b99703adc8a01cb4))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.3.0 ([b9053e4](https://github.com/offsoc/plugNmeet-server/commit/b9053e4245d73d1fd01e0f407eae67250cc87b5f))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.3.0 ([8a9ba58](https://github.com/offsoc/plugNmeet-server/commit/8a9ba58ee49519033ab750a872dc6b33f0aae135))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.3.1 ([09c7cc0](https://github.com/offsoc/plugNmeet-server/commit/09c7cc03be75c22880e32b2202622f8ef0bc6ed9))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.3.1 ([d5df05d](https://github.com/offsoc/plugNmeet-server/commit/d5df05db06f3f9644971bf8d75b90bf956fd6dd6))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.4.1 ([69b2b6c](https://github.com/offsoc/plugNmeet-server/commit/69b2b6cd9d2b5f98bc97087408c130fb44277b74))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.4.1 ([89cac3f](https://github.com/offsoc/plugNmeet-server/commit/89cac3f4eed56abf6a60ac3ef51df88b62503738))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.4.2 ([8684bff](https://github.com/offsoc/plugNmeet-server/commit/8684bffd266aa1ac2da209b42b67c8697829ce0b))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.4.2 ([d689646](https://github.com/offsoc/plugNmeet-server/commit/d6896460b5184218044ba81fd4dfc7f6b9405235))
* **deps:** update module github.com/mynaparrot/plugnmeet-protocol to v0.0.0-20240510063558-304155e97aa4 ([da59fa4](https://github.com/offsoc/plugNmeet-server/commit/da59fa4be1633be823ba5aa59b5a6ee3787bf8cb))
* **deps:** update module github.com/mynaparrot/plugnmeet-protocol to v0.0.0-20240510063558-304155e97aa4 ([ef48d4a](https://github.com/offsoc/plugNmeet-server/commit/ef48d4a7f41ce1819240b8c6e6c7619e65c1f8d4))
* **deps:** update module github.com/mynaparrot/plugnmeet-protocol to v0.0.0-20240514095941-b542c3f97e6d ([736cf65](https://github.com/offsoc/plugNmeet-server/commit/736cf6538ef28942c92499ae8a003279e42dcf87))
* **deps:** update module github.com/mynaparrot/plugnmeet-protocol to v0.0.0-20240514095941-b542c3f97e6d ([460121f](https://github.com/offsoc/plugNmeet-server/commit/460121f9ffc63b36269ba13ea1af9d5b743cc222))
* **deps:** update module github.com/nats-io/nkeys to v0.4.8 ([a3f5869](https://github.com/offsoc/plugNmeet-server/commit/a3f58690d945caaf2d01341df1a58392e4438e90))
* **deps:** update module github.com/nats-io/nkeys to v0.4.8 ([6b213d9](https://github.com/offsoc/plugNmeet-server/commit/6b213d95a7abacff4092e25fd3a9e4f8f2a3377e))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.0 ([a6ce6d5](https://github.com/offsoc/plugNmeet-server/commit/a6ce6d5e4dc196dd881a0042f8f1935e18d2e662))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.0 ([8d7dc9a](https://github.com/offsoc/plugNmeet-server/commit/8d7dc9ad888af695cdcab2562e881d76b73b0a15))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.1 ([85b851b](https://github.com/offsoc/plugNmeet-server/commit/85b851baf731dbb9ac672750db36bdf59f2350a5))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.1 ([be23bbe](https://github.com/offsoc/plugNmeet-server/commit/be23bbe3e8efd00e5a20e00b2354cb6fc011a2f2))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.3 [security] ([6f845c7](https://github.com/offsoc/plugNmeet-server/commit/6f845c732f077f72629eeb20e9b980b72d2397e3))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.3 [security] ([607dac6](https://github.com/offsoc/plugNmeet-server/commit/607dac6f292b4ac644f9ae5af6cb645f9c362c9f))
* **deps:** update module github.com/urfave/cli/v2 to v2.27.2 ([fc1de52](https://github.com/offsoc/plugNmeet-server/commit/fc1de52425fe3e99ecf2fd4eb11f67775ed8511a))
* **deps:** update module github.com/urfave/cli/v2 to v2.27.2 ([ffcae30](https://github.com/offsoc/plugNmeet-server/commit/ffcae30a2b74f2932c9699aecc95f0195e213e12))
* **deps:** update module github.com/urfave/cli/v2 to v2.27.3 ([ed32203](https://github.com/offsoc/plugNmeet-server/commit/ed3220334c433fe4ae5190626982a12554ed9b63))
* **deps:** update module github.com/urfave/cli/v2 to v2.27.3 ([2a090de](https://github.com/offsoc/plugNmeet-server/commit/2a090de3b9e482dff0f1461518f69e227684eda3))
* **deps:** update module github.com/urfave/cli/v2 to v2.27.5 ([701b06c](https://github.com/offsoc/plugNmeet-server/commit/701b06c313f2d7d994892d46a165cdd0d67a21ac))
* **deps:** update module github.com/urfave/cli/v2 to v2.27.5 ([664e915](https://github.com/offsoc/plugNmeet-server/commit/664e9153def94cddcb75c3b1ddcb6f4c810295a9))
* **deps:** upgrade `nats.go` ([bc3a121](https://github.com/offsoc/plugNmeet-server/commit/bc3a121a67920d5c082d6d6f6c4d2f570258f9bb))
* display runtime version too ([7ac5d23](https://github.com/offsoc/plugNmeet-server/commit/7ac5d23a29f686d969905b382a0803e5b77ede36))
* evaluate response from recorder ([7b81c4a](https://github.com/offsoc/plugNmeet-server/commit/7b81c4a6bba042d2887fbdd1a872349a46850527))
* few more improvements + clean up ([689d676](https://github.com/offsoc/plugNmeet-server/commit/689d6762807e7e356da69c3a1d3b3df14578e882))
* few optimization and code clean up ([bb7eb7e](https://github.com/offsoc/plugNmeet-server/commit/bb7eb7ebfa7542d4a1bcaba03f75f2c14df6422d))
* few updates ([3944420](https://github.com/offsoc/plugNmeet-server/commit/394442074371516ac16ccef00a4ad70aa6d08d72))
* **improve:** added better room creation lock system ([4353828](https://github.com/offsoc/plugNmeet-server/commit/4353828429cc23fde29fa66c6b4315c92096be9f))
* **improve:** added few more tests ([4893ebe](https://github.com/offsoc/plugNmeet-server/commit/4893ebe6c701028a203e2d64c24f2abe78564952))
* **improve:** added nats data cache feature ([7760b08](https://github.com/offsoc/plugNmeet-server/commit/7760b08fde40035fd0c6423c6c3cbea8a5f8a84b))
* **improve:** added nats data cache feature ([5502704](https://github.com/offsoc/plugNmeet-server/commit/5502704e69d64f5475a3c80a5475b19b80f1e6e8))
* **improve:** handle time way more efficiently ([c379ec2](https://github.com/offsoc/plugNmeet-server/commit/c379ec254cb0acf025eb1d7e32e71e0aa44b3601))
* **improve:** improve webhook notifier to handle each room's request in separate queue ([762809c](https://github.com/offsoc/plugNmeet-server/commit/762809c4db604251d62c5dc2bac39d6934562d80))
* **improve:** move `os.Signal` in bootup and remove unnecessary goroutine ([4eeed0c](https://github.com/offsoc/plugNmeet-server/commit/4eeed0cfedd7f18307b807066a92e1831e3e7be6))
* **improve:** optimized database query and added index ([d9c1354](https://github.com/offsoc/plugNmeet-server/commit/d9c135411b4af76bad8d86f4f05708e568d58014))
* **improve:** various improvement regarding file handling ([d4453fe](https://github.com/offsoc/plugNmeet-server/commit/d4453fe7535fd1604591dfa5256a7800b0d9cb0a))
* invalid key value used ([cdd2584](https://github.com/offsoc/plugNmeet-server/commit/cdd2584cef2cce5800cc52629e02173245b154ef))
* just for safety ([17ba8f4](https://github.com/offsoc/plugNmeet-server/commit/17ba8f4eafbfc40982d533743a8f6285176e03bc))
* missed to add `Replicas` ([fb6cba0](https://github.com/offsoc/plugNmeet-server/commit/fb6cba0c453c61797578662b56c2254263a52f5e))
* missing EmptyTimeout & MaxParticipants in webhook ([879d5b1](https://github.com/offsoc/plugNmeet-server/commit/879d5b1591553fc3c43f9f11b5ac3a8ad2d618d2))
* mistakenly removed post method ([13a6994](https://github.com/offsoc/plugNmeet-server/commit/13a69947066217a10ce9f58d82aa3cfd1f926fdb))
* moved poll to its own features option with backward compatibility ([b361af4](https://github.com/offsoc/plugNmeet-server/commit/b361af4161d6b134be037084c86f8bbc3c252e1f))
* moved to proto for reusing ([517d3ad](https://github.com/offsoc/plugNmeet-server/commit/517d3ad38c3067964fda5318e9517a7380e4dfc7))
* multiuser whiteboard sync problem ([534c196](https://github.com/offsoc/plugNmeet-server/commit/534c19616d1d0d631c470d009033eb801e73792f))
* nil checker to avoid `nil pointer dereference` error ([0f5fc9d](https://github.com/offsoc/plugNmeet-server/commit/0f5fc9d25e911ccbed9394f8387ac620a5abf885))
* option to use nkey ([9bd1964](https://github.com/offsoc/plugNmeet-server/commit/9bd196401a8bb09c792f4b93496564a7404682d3))
* optional for encrypting the nats request payloads ([4bfd3e9](https://github.com/offsoc/plugNmeet-server/commit/4bfd3e9d72982eb21f633a8682ac1056ee70125e))
* prevent using reserved name ([e5b3b13](https://github.com/offsoc/plugNmeet-server/commit/e5b3b13ec04e5d70a399ac326d81d32fdbfad434))
* protocol upgrade ([5f7185f](https://github.com/offsoc/plugNmeet-server/commit/5f7185fd502f0f3cf651750571a7d7ffea353d1e))
* recording server will authenticate with plugNmeet server for nats ([dee6082](https://github.com/offsoc/plugNmeet-server/commit/dee608242b7590e09d4a4c9a6515eb19e8c1167f))
* release-please-action ([3223400](https://github.com/offsoc/plugNmeet-server/commit/3223400ec426f66fd8bbf2628877e6801105be37))
* remove duplicate tag ([377903f](https://github.com/offsoc/plugNmeet-server/commit/377903f6347033da290c381c12540021259d201c))
* replacing with nkey for new installations ([30e1463](https://github.com/offsoc/plugNmeet-server/commit/30e146336d63c78b85f043caabb6f1be8528a708))
* showing invalid poll result ([21ed66d](https://github.com/offsoc/plugNmeet-server/commit/21ed66dc367a76f592815718d6a545777ccec85d))
* some clean up ([2472ad3](https://github.com/offsoc/plugNmeet-server/commit/2472ad3ccabda7a7de0006a5deec60996675342f))
* time zone problem ([5b655f5](https://github.com/offsoc/plugNmeet-server/commit/5b655f514bfa7debe88fdf883aea612e95880cca))
* updated text ([4081b07](https://github.com/offsoc/plugNmeet-server/commit/4081b0751d680fe3e3c8cab50922f45873fc34c0))
* **update:** few clean ups ([3d632ec](https://github.com/offsoc/plugNmeet-server/commit/3d632ecd5f4394d5759b2ab8c89facf3ffeb5142))
* upgrade urfave/cli to `v3` ([887bc76](https://github.com/offsoc/plugNmeet-server/commit/887bc76ebb8a5df0410af1286772635e489613ea))
* use `StatusServiceUnavailable` ([554e1f5](https://github.com/offsoc/plugNmeet-server/commit/554e1f5e8167ccf2bb91b796f5ffc2369f598908))
* use proper format ([a20d0b6](https://github.com/offsoc/plugNmeet-server/commit/a20d0b6b99016483d5ebf32eb96a9cb2a8066cad))
* user appropriate library where necessary ([dcb7e6e](https://github.com/offsoc/plugNmeet-server/commit/dcb7e6e4c866ffab512e8fef9afab7e9fb5205b6))

## [1.9.1](https://github.com/mynaparrot/plugNmeet-server/compare/v1.9.0...v1.9.1) (2025-08-12)


### Bug Fixes

* nil checker to avoid `nil pointer dereference` error ([0f5fc9d](https://github.com/mynaparrot/plugNmeet-server/commit/0f5fc9d25e911ccbed9394f8387ac620a5abf885))
* protocol upgrade ([5f7185f](https://github.com/mynaparrot/plugNmeet-server/commit/5f7185fd502f0f3cf651750571a7d7ffea353d1e))

## [1.9.0](https://github.com/mynaparrot/plugNmeet-server/compare/v1.8.4...v1.9.0) (2025-07-15)


### Features

* dependencies injection using wire ([0d6ddef](https://github.com/mynaparrot/plugNmeet-server/commit/0d6ddefc6b718d7c546ab0b582c21ed46062cf02))
* dependencies injection using wire ([a9e78ff](https://github.com/mynaparrot/plugNmeet-server/commit/a9e78ff03942386ce3d6177f8a7e8aabbd2f0a17))


### Bug Fixes

* cleaned up logic ([4aa5b49](https://github.com/mynaparrot/plugNmeet-server/commit/4aa5b4930286af2d5379cafc3b914bcf36cde083))
* **cleanup:** remove unnecessary goroutine ([b5f6d0a](https://github.com/mynaparrot/plugNmeet-server/commit/b5f6d0a82a232ad3d97ef82ffe4bfe47a86407ac))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to b2a5cc6 ([ce7e129](https://github.com/mynaparrot/plugNmeet-server/commit/ce7e12944665765bc7918d6fe1d5425bd413748a))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to b2a5cc6 ([3e40713](https://github.com/mynaparrot/plugNmeet-server/commit/3e4071398f9c55b6e305a242f789f0d53126e3e4))
* **deps:** update module github.com/ansrivas/fiberprometheus/v2 to v2.12.0 ([b5c2b26](https://github.com/mynaparrot/plugNmeet-server/commit/b5c2b26cc02fc0be77bbdf660350fccb9a0ab04e))
* **deps:** update module github.com/ansrivas/fiberprometheus/v2 to v2.12.0 ([3bda07e](https://github.com/mynaparrot/plugNmeet-server/commit/3bda07e3d1b7abd4f863547e21aca9a2245d03c6))
* **improve:** handle time way more efficiently ([c379ec2](https://github.com/mynaparrot/plugNmeet-server/commit/c379ec254cb0acf025eb1d7e32e71e0aa44b3601))
* **improve:** improve webhook notifier to handle each room's request in separate queue ([762809c](https://github.com/mynaparrot/plugNmeet-server/commit/762809c4db604251d62c5dc2bac39d6934562d80))
* **improve:** move `os.Signal` in bootup and remove unnecessary goroutine ([4eeed0c](https://github.com/mynaparrot/plugNmeet-server/commit/4eeed0cfedd7f18307b807066a92e1831e3e7be6))
* **improve:** optimized database query and added index ([d9c1354](https://github.com/mynaparrot/plugNmeet-server/commit/d9c135411b4af76bad8d86f4f05708e568d58014))
* **improve:** various improvement regarding file handling ([d4453fe](https://github.com/mynaparrot/plugNmeet-server/commit/d4453fe7535fd1604591dfa5256a7800b0d9cb0a))
* mistakenly removed post method ([13a6994](https://github.com/mynaparrot/plugNmeet-server/commit/13a69947066217a10ce9f58d82aa3cfd1f926fdb))

## [1.8.4](https://github.com/mynaparrot/plugNmeet-server/compare/v1.8.3...v1.8.4) (2025-06-27)


### Bug Fixes

* bug + clean up ([45bca50](https://github.com/mynaparrot/plugNmeet-server/commit/45bca50d19638a3f7e273d2a38f6d7fdd05f8bbc))
* bump proto ([acb2bbf](https://github.com/mynaparrot/plugNmeet-server/commit/acb2bbf9a0d80041483313c3404189a813795f1e))
* **deps:** protocol update ([a321514](https://github.com/mynaparrot/plugNmeet-server/commit/a32151457d29ff33ca8fd5be7a27f9f12a731c50))
* **deps:** update module github.com/ansrivas/fiberprometheus/v2 to v2.10.0 ([7364863](https://github.com/mynaparrot/plugNmeet-server/commit/7364863bad7628530fce3f17a8a40fca038a6ce0))
* **deps:** update module github.com/ansrivas/fiberprometheus/v2 to v2.10.0 ([cfbb9f3](https://github.com/mynaparrot/plugNmeet-server/commit/cfbb9f336e38382a4802a700b194f0cc92b5e81c))
* few more improvements + clean up ([689d676](https://github.com/mynaparrot/plugNmeet-server/commit/689d6762807e7e356da69c3a1d3b3df14578e882))
* few optimization and code clean up ([bb7eb7e](https://github.com/mynaparrot/plugNmeet-server/commit/bb7eb7ebfa7542d4a1bcaba03f75f2c14df6422d))
* **improve:** added better room creation lock system ([4353828](https://github.com/mynaparrot/plugNmeet-server/commit/4353828429cc23fde29fa66c6b4315c92096be9f))
* **improve:** added few more tests ([4893ebe](https://github.com/mynaparrot/plugNmeet-server/commit/4893ebe6c701028a203e2d64c24f2abe78564952))
* **improve:** added nats data cache feature ([7760b08](https://github.com/mynaparrot/plugNmeet-server/commit/7760b08fde40035fd0c6423c6c3cbea8a5f8a84b))
* **improve:** added nats data cache feature ([5502704](https://github.com/mynaparrot/plugNmeet-server/commit/5502704e69d64f5475a3c80a5475b19b80f1e6e8))
* updated text ([4081b07](https://github.com/mynaparrot/plugNmeet-server/commit/4081b0751d680fe3e3c8cab50922f45873fc34c0))

## [1.8.3](https://github.com/mynaparrot/plugNmeet-server/compare/v1.8.2...v1.8.3) (2025-05-05)


### Bug Fixes

* **bug:** ingress was removing because of invalid inactivity checking ([c08ec4a](https://github.com/mynaparrot/plugNmeet-server/commit/c08ec4ab3f060183c328d0d40a692ab8ab455ce6))
* **bug:** user's online status wasn't updated properly in the latest version of nats-server ([f9d9793](https://github.com/mynaparrot/plugNmeet-server/commit/f9d97932e902832455f7f4ce6e13bd8f702679ab))
* **deps:** deps update ([90d40b3](https://github.com/mynaparrot/plugNmeet-server/commit/90d40b3a6faa84da20889687d286fa62acb13b0e))
* **deps:** deps update ([a03f37e](https://github.com/mynaparrot/plugNmeet-server/commit/a03f37e3cc9efa3baf7c7866157460c27350c1b7))
* user appropriate library where necessary ([dcb7e6e](https://github.com/mynaparrot/plugNmeet-server/commit/dcb7e6e4c866ffab512e8fef9afab7e9fb5205b6))

## [1.8.2](https://github.com/mynaparrot/plugNmeet-server/compare/v1.8.1...v1.8.2) (2025-04-10)


### Bug Fixes

* added option to insert a self provided e2ee key ([386cbdd](https://github.com/mynaparrot/plugNmeet-server/commit/386cbdd1ce79e32ce206f18164fc7ed84abffc64))
* api to upload base64 encoded files ([f7b9923](https://github.com/mynaparrot/plugNmeet-server/commit/f7b992311ab4ef0a0214030444bb8fd1b420ca43))
* bump proto ([21b0c2d](https://github.com/mynaparrot/plugNmeet-server/commit/21b0c2de01356ad588d53244c51e17d788d98214))
* deps update ([264d0ee](https://github.com/mynaparrot/plugNmeet-server/commit/264d0eecf750ece8ebc77cd8361997003e4df019))
* **deps:** update module github.com/go-jose/go-jose/v4 to v4.0.5 [security] ([42be81b](https://github.com/mynaparrot/plugNmeet-server/commit/42be81ba43c996af7d0ee9bcb4cfe79b1c11fd21))
* **deps:** update module github.com/go-jose/go-jose/v4 to v4.0.5 [security] ([084bc36](https://github.com/mynaparrot/plugNmeet-server/commit/084bc367dbf9ac2fa021ccf0fb42caa651cbd01d))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.1 ([85b851b](https://github.com/mynaparrot/plugNmeet-server/commit/85b851baf731dbb9ac672750db36bdf59f2350a5))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.1 ([be23bbe](https://github.com/mynaparrot/plugNmeet-server/commit/be23bbe3e8efd00e5a20e00b2354cb6fc011a2f2))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.3 [security] ([6f845c7](https://github.com/mynaparrot/plugNmeet-server/commit/6f845c732f077f72629eeb20e9b980b72d2397e3))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.3 [security] ([607dac6](https://github.com/mynaparrot/plugNmeet-server/commit/607dac6f292b4ac644f9ae5af6cb645f9c362c9f))
* display runtime version too ([7ac5d23](https://github.com/mynaparrot/plugNmeet-server/commit/7ac5d23a29f686d969905b382a0803e5b77ede36))
* moved poll to its own features option with backward compatibility ([b361af4](https://github.com/mynaparrot/plugNmeet-server/commit/b361af4161d6b134be037084c86f8bbc3c252e1f))
* upgrade urfave/cli to `v3` ([887bc76](https://github.com/mynaparrot/plugNmeet-server/commit/887bc76ebb8a5df0410af1286772635e489613ea))

## [1.8.1](https://github.com/mynaparrot/plugNmeet-server/compare/v1.8.0...v1.8.1) (2025-02-20)


### Bug Fixes

* bump proto ([9fcde46](https://github.com/mynaparrot/plugNmeet-server/commit/9fcde46144b7a0511dc77d218276b5f32f99e97b))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.9.2 ([e924e7b](https://github.com/mynaparrot/plugNmeet-server/commit/e924e7b1f0bb481b367b62f794438b1e3b5098c1))
* just for safety ([17ba8f4](https://github.com/mynaparrot/plugNmeet-server/commit/17ba8f4eafbfc40982d533743a8f6285176e03bc))
* moved to proto for reusing ([517d3ad](https://github.com/mynaparrot/plugNmeet-server/commit/517d3ad38c3067964fda5318e9517a7380e4dfc7))
* option to use nkey ([9bd1964](https://github.com/mynaparrot/plugNmeet-server/commit/9bd196401a8bb09c792f4b93496564a7404682d3))
* optional for encrypting the nats request payloads ([4bfd3e9](https://github.com/mynaparrot/plugNmeet-server/commit/4bfd3e9d72982eb21f633a8682ac1056ee70125e))
* prevent using reserved name ([e5b3b13](https://github.com/mynaparrot/plugNmeet-server/commit/e5b3b13ec04e5d70a399ac326d81d32fdbfad434))
* recording server will authenticate with plugNmeet server for nats ([dee6082](https://github.com/mynaparrot/plugNmeet-server/commit/dee608242b7590e09d4a4c9a6515eb19e8c1167f))
* replacing with nkey for new installations ([30e1463](https://github.com/mynaparrot/plugNmeet-server/commit/30e146336d63c78b85f043caabb6f1be8528a708))
* some clean up ([2472ad3](https://github.com/mynaparrot/plugNmeet-server/commit/2472ad3ccabda7a7de0006a5deec60996675342f))
* use proper format ([a20d0b6](https://github.com/mynaparrot/plugNmeet-server/commit/a20d0b6b99016483d5ebf32eb96a9cb2a8066cad))

## [1.8.0](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.8...v1.8.0) (2025-02-01)


### Features

* option to keep deleted recordings in backup ([7d1ee5d](https://github.com/mynaparrot/plugNmeet-server/commit/7d1ee5db73d3689d7137233ba0dbe48294eb82a6))


### Bug Fixes

* deps update ([93fff85](https://github.com/mynaparrot/plugNmeet-server/commit/93fff85fe110fc79da5a602c51d2f91c8031e321))
* **deps:** update module github.com/goccy/go-json to v0.10.5 ([93ef56e](https://github.com/mynaparrot/plugNmeet-server/commit/93ef56eeeb28292a1d0409684bc71a624bd414aa))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.4.2 ([d689646](https://github.com/mynaparrot/plugNmeet-server/commit/d6896460b5184218044ba81fd4dfc7f6b9405235))
* few updates ([3944420](https://github.com/mynaparrot/plugNmeet-server/commit/394442074371516ac16ccef00a4ad70aa6d08d72))
* remove duplicate tag ([377903f](https://github.com/mynaparrot/plugNmeet-server/commit/377903f6347033da290c381c12540021259d201c))

## [1.7.8](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.7...v1.7.8) (2025-01-18)


### Bug Fixes

* bump nodejs ([2779eab](https://github.com/mynaparrot/plugNmeet-server/commit/2779eab3538a008d693e9a577b07aa78c5640a0e))
* **ci:** added beta release ([e7497e2](https://github.com/mynaparrot/plugNmeet-server/commit/e7497e2a84f3768007ecb1a787feea0ee87b883c))
* **ci:** always-update ([1134c50](https://github.com/mynaparrot/plugNmeet-server/commit/1134c50f91c84c7fc9eabf5e82f9b50ffdc01472))
* **ci:** bump beta version ([b79a0ea](https://github.com/mynaparrot/plugNmeet-server/commit/b79a0eadd93a0e06070f12dc087c4f18ec49247a))
* **ci:** bump beta version ([24c9082](https://github.com/mynaparrot/plugNmeet-server/commit/24c90821b4e9309ec6666ddcc19caf0f6069168d))
* **ci:** removed rebase again ([db27ad5](https://github.com/mynaparrot/plugNmeet-server/commit/db27ad54d16a4b39c9b5347ee2f06a8164a95266))
* **ci:** To rebase again ([e1681fc](https://github.com/mynaparrot/plugNmeet-server/commit/e1681fc79aede64a5a9cc297d7a4505461767eb8))
* **ci:** To rebase again ([161b04e](https://github.com/mynaparrot/plugNmeet-server/commit/161b04ed775d25a873617c229a0a8bfb23989672))
* **ci:** To rebase again ([4a34670](https://github.com/mynaparrot/plugNmeet-server/commit/4a3467056d97d57245219b0f7cb6d406c3a1fd92))
* deps update ([0d58c13](https://github.com/mynaparrot/plugNmeet-server/commit/0d58c1323a618e4fe6c2bdcb2fd853409851e3a5))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.8 ([703623d](https://github.com/mynaparrot/plugNmeet-server/commit/703623d55bbcd8a2b9c613dbe07446670e3f2ccb))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.4.1 ([89cac3f](https://github.com/mynaparrot/plugNmeet-server/commit/89cac3f4eed56abf6a60ac3ef51df88b62503738))
* evaluate response from recorder ([7b81c4a](https://github.com/mynaparrot/plugNmeet-server/commit/7b81c4a6bba042d2887fbdd1a872349a46850527))
* use `StatusServiceUnavailable` ([554e1f5](https://github.com/mynaparrot/plugNmeet-server/commit/554e1f5e8167ccf2bb91b796f5ffc2369f598908))

## [1.7.7](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.6...v1.7.7) (2024-12-20)


### Bug Fixes

* **bug:** Ingress was not working in the new Nats solution. Ref: https://github.com/mynaparrot/plugNmeet-server/discussions/611 ([2fc55a4](https://github.com/mynaparrot/plugNmeet-server/commit/2fc55a4e5c02787962a4e77804cfb83b38c7d975))
* **deps:** deps update ([3d8d4b9](https://github.com/mynaparrot/plugNmeet-server/commit/3d8d4b933cdb6fd8a992ddc672d22c83795a2be6))
* **deps:** upgrade `nats.go` ([bc3a121](https://github.com/mynaparrot/plugNmeet-server/commit/bc3a121a67920d5c082d6d6f6c4d2f570258f9bb))
* **update:** few clean ups ([3d632ec](https://github.com/mynaparrot/plugNmeet-server/commit/3d632ecd5f4394d5759b2ab8c89facf3ffeb5142))

## [1.7.6](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.5...v1.7.6) (2024-12-07)


### Bug Fixes

* **deps:** dependencies update ([ef83e0d](https://github.com/mynaparrot/plugNmeet-server/commit/ef83e0daaecb76dc884f53c4c5187d65619d3968))
* **deps:** update module github.com/nats-io/nkeys to v0.4.8 ([6b213d9](https://github.com/mynaparrot/plugNmeet-server/commit/6b213d95a7abacff4092e25fd3a9e4f8f2a3377e))
* missed to add `Replicas` ([fb6cba0](https://github.com/mynaparrot/plugNmeet-server/commit/fb6cba0c453c61797578662b56c2254263a52f5e))

## [1.7.5](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.4...v1.7.5) (2024-11-25)


### Bug Fixes

* added option `-trimpath` ([6b355d9](https://github.com/mynaparrot/plugNmeet-server/commit/6b355d9a89c668e400fefbcfea4dcfae54e14bee))
* **bump:** bump version ([de8f4dd](https://github.com/mynaparrot/plugNmeet-server/commit/de8f4dd1f0dd5086684738c7d762360cfccd28d8))
* clean up ([6e12ff1](https://github.com/mynaparrot/plugNmeet-server/commit/6e12ff16e7e5dea4e2c41c64dee5e1e8cdb13053))
* **deps:** update github.com/mynaparrot/plugnmeet-protocol digest to 4d9a8f6 ([d9ec4b4](https://github.com/mynaparrot/plugNmeet-server/commit/d9ec4b43c9abebe5ee05a59ffc9e01b417090691))
* **deps:** update module github.com/bufbuild/protovalidate-go to v0.7.3 ([0c51ba8](https://github.com/mynaparrot/plugNmeet-server/commit/0c51ba89ad3037b71bf5a3f7ae83272811c0a78e))
* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.7 ([56f5d3b](https://github.com/mynaparrot/plugNmeet-server/commit/56f5d3becfacba5c535cf45615c2fdeaf898274d))
* **deps:** update module github.com/livekit/protocol to v1.27.1 ([373f1ba](https://github.com/mynaparrot/plugNmeet-server/commit/373f1ba67264263816b1fb0a912f935d44781e5e))
* **deps:** update module github.com/livekit/server-sdk-go/v2 to v2.3.1 ([d5df05d](https://github.com/mynaparrot/plugNmeet-server/commit/d5df05db06f3f9644971bf8d75b90bf956fd6dd6))

## [1.7.4](https://github.com/mynaparrot/plugNmeet-server/compare/v1.7.3...v1.7.4) (2024-11-04)


### Bug Fixes

* **deps:** update module github.com/gabriel-vasile/mimetype to v1.4.6 ([fc6df57](https://github.com/mynaparrot/plugNmeet-server/commit/fc6df57253b8ff035cc87e3d76d651e2391a9419))
* **deps:** update module github.com/redis/go-redis/v9 to v9.7.0 ([8d7dc9a](https://github.com/mynaparrot/plugNmeet-server/commit/8d7dc9ad888af695cdcab2562e881d76b73b0a15))
* missing EmptyTimeout & MaxParticipants in webhook ([879d5b1](https://github.com/mynaparrot/plugNmeet-server/commit/879d5b1591553fc3c43f9f11b5ac3a8ad2d618d2))
* release-please-action ([3223400](https://github.com/mynaparrot/plugNmeet-server/commit/3223400ec426f66fd8bbf2628877e6801105be37))
