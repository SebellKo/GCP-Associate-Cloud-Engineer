# Chapter 15. Shared VPC·Peering·VPN·Interconnect

[← 이전 장](chapter-14.md) · [전체 목차](../README.md) · [다음 장 →](chapter-16.md)

공식 가이드 연결: **2.3** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 연결 요구를 먼저 분류한다
Shared VPC는 host project가 네트워크를 중앙 관리하고 service project가 그 네트워크를 사용하는 구조다. 프로젝트별 앱 권한과 중앙 네트워크 운영을 분리할 수 있다. VPC Network Peering은 서로 다른 VPC를 내부 주소로 연결하며 CIDR 겹침과 비전이성 제약을 확인한다. A-B, B-C 피어링이 A-C 통신을 자동으로 제공하지 않는다.

Cloud VPN은 IPsec 기반 터널을 사용해 온프레미스와 연결한다. HA 구성은 터널·라우터·대칭 경로 등 실제 토폴로지 조건을 만족해야 한다. Interconnect는 전용 또는 파트너 연결을 검토하는 선택지다. Interconnect라는 이유만으로 모든 트래픽이 IPsec 암호화된다고 가정하지 않는다.

연결 장애는 주소 겹침, route 교환, BGP 상태, firewall, DNS 순서로 증거를 모아 좁힌다. 네트워크 연결과 Google API IAM 권한은 별개다. 연결됐어도 접근 주체에 권한이 없으면 API는 거부한다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. Shared VPC와 peering은 같은가?

2. peering은 자동 전이적인가?

3. 사설 연결이 API 권한까지 부여하는가?

## 3. 실습과 관찰

무료 설계 실습: 세 VPC A/B/C의 CIDR을 정하고 A-B, B-C peering에서 A-C 연결이 자동 성립하지 않는 이유를 설명한다. 중앙 네트워크팀과 앱팀의 프로젝트 책임을 표로 적는다. VPN/Interconnect는 실제 계약·요금이 있어 이 교재에서는 설계와 상태 화면 읽기를 기본으로 한다. 새 터널이나 회선을 만들 필요 없다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

중앙팀이 서브넷·방화벽을 관리하고 앱팀은 각자 프로젝트에 VM을 배포한다.

A. 각 앱에 조직 Owner를 준다
B. 버킷을 공개한다
C. Shared VPC host/service project를 구성한다
D. 모든 프로젝트를 삭제해 하나로 합친다

### 문제 2 — 단일 선택

A-B와 B-C peering 후 A-C가 통신되지 않는다.

A. Storage 위치를 바꾼다
B. DNS TTL만 늘린다
C. 비전이적 peering 특성을 고려해 지원되는 별도 연결을 설계한다
D. B를 재부팅하면 항상 해결된다

### 문제 3 — 단일 선택

작은 지사의 온프레미스와 GCP를 IPsec으로 연결해야 한다.

A. Cloud Storage만 만든다
B. Cloud Identity 그룹만 만든다
C. MIG autohealing을 켠다
D. Cloud VPN을 검토한다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. 중앙 네트워크 공유와 별도 네트워크 연결이다.

2. 아니다.

3. 아니다.

### 시험형 예제

**1. C** — 관리 책임을 분리하는 Shared VPC 요구다. 프로젝트 통합·공개화·Owner는 책임 분리 요구에 맞지 않는다.

**2. C** — 피어링은 전이 연결을 자동 제공하지 않는다. 재부팅·DNS·Storage는 토폴로지 제약을 없애지 않는다.

**3. D** — IPsec 연결 요구는 VPN에 해당한다. 객체 저장·VM 교체·그룹 생성은 터널을 제공하지 않는다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/vpc/docs/shared-vpc)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-14.md) · [다음 장 →](chapter-16.md)
