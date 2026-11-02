# Chapter 14. Cloud NAT·DNS·고정 IP·로드밸런서

[← 이전 장](chapter-13.md) · [전체 목차](../README.md) · [다음 장 →](chapter-15.md)

공식 가이드 연결: **2.3 / 3.3** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 발신과 수신 경로를 분리한다
Public Cloud NAT는 외부 IP 없는 VM 등의 인터넷 발신 통신을 지원한다. 새 inbound 연결을 받는 공개 서비스 엔드포인트가 아니다. Private Google Access는 내부 IP만 있는 VM의 Google API 접근을 지원하는 별도 기능이다. NAT가 모든 API·사설 경로 문제의 만능 해결책이라고 생각하지 않는다.

Cloud DNS의 public zone은 공개 이름 해석, private zone은 허용된 VPC 범위의 내부 해석에 사용한다. 정적 IP는 주소 안정성을 제공하며 사용 범위가 regional/global인지 제품과 맞춰야 한다. DNS 변경 후에는 TTL과 캐시도 확인한다.

Application Load Balancer는 HTTP(S) 요구, Network Load Balancer는 TCP/UDP 등 L4 요구를 기준으로 고른다. internal/external, regional/global, proxy/passthrough 조건도 비교한다. Network Service Tiers의 Premium/Standard는 지원되는 서비스와 트래픽 경로·비용 차이를 확인한다. LB는 backend health와 capacity가 모두 필요하다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. NAT가 새 inbound 연결을 열어 주는가?

2. HTTP 경로별 분기라면 어떤 LB 계열인가?

3. private DNS zone만 만들면 인터넷 사용자도 해석하는가?

## 3. 실습과 관찰

구성도에서 외부 사용자→LB→VM, VM→NAT→인터넷, VM→Google API 세 경로를 따로 표시한다. 선택 유료 실습은 Cloud Router+NAT를 구성해 private VM의 외부 접속을 확인하는 것이다. 완료 후 NAT→Router를 삭제하고 정적 IP와 LB forwarding rule 잔존 여부를 검사한다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

외부 IP 없는 VM이 인터넷 업데이트 서버에 접속해야 한다.

A. private DNS zone만 만든다
B. 버킷을 Archive로 바꾼다
C. ingress 80만 연다
D. 적절한 경로·egress와 Public Cloud NAT를 구성한다

### 문제 2 — 단일 선택

서비스가 HTTP 경로 /api와 /images에 따라 다른 backend로 분기해야 한다.

A. Application Load Balancer를 검토한다
B. Cloud NAT만 사용한다
C. DNS A 레코드에 URL path를 저장한다
D. 블록 디스크를 확장한다

### 문제 3 — 단일 선택

외부 파트너가 고정 IP allowlist를 요구한다.

A. 매번 임시 IP를 재발급한다
B. 해당 통신 경로의 예약 정적 IP를 사용한다
C. IAM Viewer만 준다
D. DNS TTL만 0으로 한다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. 아니다.

2. Application Load Balancer다.

3. 아니다.

### 시험형 예제

**1. D** — 발신 주소 변환과 경로가 필요하다. DNS·inbound HTTP·객체 등급은 인터넷 발신을 제공하지 않는다.

**2. A** — L7 URL 기반 분기는 Application LB의 영역이다. NAT·디스크·A 레코드는 경로 라우팅을 하지 않는다.

**3. B** — 안정된 IP가 요구다. 임시 주소·TTL·IAM 역할은 고정 발신 또는 수신 주소를 보장하지 않는다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/nat/docs/overview)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-13.md) · [다음 장 →](chapter-15.md)
