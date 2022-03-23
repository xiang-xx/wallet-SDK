VERSION=`git describe --tags --dirty`
DATE=`date +%FT%T%z`

outdir=out

module=github.com/coming-chat/wallet-SDK

pkgCore = ${module}/core/eth

pkgEth =  $(pkgCore)

pkgPolka = ${module}/wallet

pkgSDK = ${module}/sdk/core

buildAllSDKAndroid:
	gomobile bind -ldflags "-s -w" -target=android -o=${outdir}/wallet.aar ${pkgEth} ${pkgPolka}

buildAllSDKIOS:
	gomobile bind -ldflags "-s -w" -target=ios  -o=${outdir}/Wallet.xcframework ${pkgEth} ${pkgPolka}

buildSDK:
	gomobile bind -ldflags "-s -w" -target=ios  -o=${outdir}/SDK.xcframework ${pkgSDK}

packageAll:
	rm -rf ${outdir}/*
	@make buildAllAndroid && make buildAllIOS
	@cd ${outdir} && mkdir android && mv wallet* android
	@cd ${outdir} && tar czvf android.tar.gz android/*
	@cd ${outdir} && tar czvf Wallet.xcframework Wallet.xcframework/*