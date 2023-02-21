package util

type ClientStateEnum int

const (
	_ ClientStateEnum = iota
	ClientStateNotVerify
	ClientStateVerified
	ClientStateNeedCharacter
	ClientStateInited
	ClientStateMatchingWaiting      //다른서버로 갈꺼야 매칭중
	ClientStateConnectedOtherServer //다른서버로 가야함
	ClientStateErrorKickOut
	ClientStateReconnectWaiting     //재연결 대기중 move hide
	ClientStateReconnectedClose     //재연결되어 필요없는 소켓?
	ClientStateOverlappedLoginClose //중복 로그인이 되서 종료?
	ClientStateClosed
	ClientStateContentsReEnterWaiting //컨텐츠 재진입 대기중?
)
