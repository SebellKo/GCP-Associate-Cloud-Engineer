# Chapter 13. VPC·서브넷·라우팅·방화벽

[← 이전 장](chapter-12.md) · [전체 목차](../README.md) · [다음 장 →](chapter-14.md)

공식 가이드 연결: **2.3 / 3.3** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 네트워크의 범위를 이해한다
Google Cloud VPC는 global, subnet은 regional, VM은 zonal 자원이다. 같은 리전의 여러 존 VM이 같은 서브넷을 쓸 수 있다. custom mode VPC는 서브넷 CIDR을 직접 설계한다. 주소 범위는 다른 연결 네트워크와 겹치지 않게 계획한다. 서브넷 IPv4 범위는 제약 안에서 확장할 수 있지만 축소를 동일하게 생각하지 않는다.

Route는 패킷을 어디로 전달할지, firewall은 통신을 허용할지 정한다. 기본 implied ingress deny·egress allow와 default network가 추가하는 명시 규칙을 구분한다. 낮은 priority 숫자가 더 높은 우선순위다. 방향·source·destination·target·protocol·port를 모두 읽는다.

전통적 network tag와 IAM으로 관리하는 secure Tag는 다르다. Cloud NGFW 정책은 secure Tags나 service account 등 지원되는 속성으로 대상을 지정한다. 정책 계층·평가 순서를 확인하지 않고 개별 VPC 규칙 하나만 보고 최종 허용을 단정하지 않는다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. VPC/subnet/VM의 범위는?

2. 경로가 있다는 것과 허용됐다는 것은 같은가?

3. priority 100과 1000 중 먼저 평가되는 숫자는?

## 3. 실습과 관찰

[VM 실습](../labs/vm.md)의 custom VPC와 subnet을 조회한다. `gcloud compute networks list`, `gcloud compute networks subnets list`, `gcloud compute firewall-rules list`의 범위를 비교한다. ingress source가 IAP 범위로 제한됐는지 확인한다. 생성했다면 VM을 먼저 삭제한 뒤 규칙·서브넷·VPC를 삭제한다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

같은 리전의 두 존 VM에 동일 서브넷을 쓰고 싶다.

A. Cloud DNS가 서브넷을 대신한다
B. regional subnet을 두 존에서 사용한다
C. 존마다 반드시 VPC를 만든다
D. 서브넷을 global로 바꾼다

### 문제 2 — 단일 선택

앱 서버만 DB TCP 포트에 접근하도록 제한한다.

A. 방화벽을 삭제한다
B. DB 대상을 제한하고 앱의 허용 source와 필요한 포트만 연다
C. DB를 public bucket으로 바꾼다
D. 0.0.0.0/0의 모든 포트를 허용한다

### 문제 3 — 단일 선택

서브넷 IP가 부족하다. 기존 VM 주소를 유지하며 대응하려 한다.

A. Storage 등급을 바꾼다
B. 겹침과 제약을 확인한 뒤 primary IPv4 범위 확장을 검토한다
C. DNS 이름 길이를 줄인다
D. 반드시 모든 VM을 삭제한다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. global/regional/zonal이다.

2. 아니다.

3. 100이다.

### 시험형 예제

**1. B** — 서브넷은 regional이므로 같은 리전 여러 존에서 쓴다. VPC 분리 강제나 global 전환은 잘못된 전제다.

**2. B** — source·target·port를 좁혀 최소 접근을 구현한다. 전체 허용·공개화·삭제는 요구를 깨뜨린다.

**3. B** — 지원되는 범위 확장이 직접적이다. VM 삭제는 항상 필수가 아니고 DNS·스토리지는 IP 수를 늘리지 않는다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/vpc/docs/vpc)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-12.md) · [다음 장 →](chapter-14.md)
