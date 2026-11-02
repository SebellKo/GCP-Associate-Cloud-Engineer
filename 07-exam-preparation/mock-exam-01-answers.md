# 모의고사 1 정답·해설

[문제](mock-exam-01.md)

정답 집합 전체 일치로 채점한다. 같은 개념을 틀리면 해당 장의 실습·선택 기준으로 돌아간다.

## Q01. A

정책은 상위에서 일관되게 적용하고 프로젝트로 경계를 분리한다. VM 이름·Owner·공개화는 경계를 만들지 않는다.

복습: [Chapter 03](../chapters/chapter-03.md) · 영역 1

## Q02. C

그룹 기반 권한 관리가 반복 바인딩을 줄인다. 비밀 공유는 identity를 잃게 하고 VM 추가는 무관하다.

복습: [Chapter 03](../chapters/chapter-03.md) · 영역 1

## Q03. A

API 사용 준비 문제를 해결해야 한다. 네트워크 삭제·공개화·DNS는 API 활성화가 아니다.

복습: [Chapter 05](../chapters/chapter-05.md) · 영역 1

## Q04. A

같은 리전 quota는 다른 존에서도 공유될 수 있다. 증액은 승인·capacity 보장이 아니며 나머지는 한도를 늘리지 않는다.

복습: [Chapter 05](../chapters/chapter-05.md) · 영역 1

## Q05. D

서비스별 위치·기능 가용성이 다르다. 나머지는 위치 지원의 근거가 아니다.

복습: [Chapter 05](../chapters/chapter-05.md) · 영역 1

## Q06. C

실행 대상을 먼저 바로잡는다. 적용·로그 삭제·권한 확대는 잘못된 대상 문제를 키운다.

복습: [Chapter 04](../chapters/chapter-04.md) · 영역 1

## Q07. C

자원·정책 인벤토리 검색 목적이다. 결제 통화·DNS·휴지통은 자산 목록을 제공하지 않는다.

복습: [Chapter 05](../chapters/chapter-05.md) · 영역 1

## Q08. C

결제 권한은 리소스 운영 권한과 별도다. 디스크·키 공개·Spot은 결제 연결 권한이 아니다.

복습: [Chapter 06](../chapters/chapter-06.md) · 영역 1

## Q09. C

결제 export는 분석 가능한 비용 데이터를 제공한다. 로그만으로 전체 비용을 보장하지 못하고 NAT·이미지는 무관하다.

복습: [Chapter 06](../chapters/chapter-06.md) · 영역 1

## Q10. B, C

알림과 차단을 구분하고 비용 집계 지연을 고려한다. 나머지는 과금·자원 수명에 대한 잘못된 가정이다.

복습: [Chapter 06](../chapters/chapter-06.md) · 영역 1

## Q11. D

사람의 외부 identity 연결은 Workforce 용도다. 앱 키 공유·NAT·MIG는 직원 SSO 구성이 아니다.

복습: [Chapter 25](../chapters/chapter-25.md) · 영역 1

## Q12. A

AI 제안도 실제 환경의 요구와 맞는지 확인한다. 무검토 적용·관측 제거·과도 권한은 근거를 제공하지 않는다.

복습: [Chapter 05](../chapters/chapter-05.md) · 영역 1

## Q13. B

OS 제어 요구에 VM이 적합하다. 다른 선택은 해당 실행 제어를 제공하지 않는다.

복습: [Chapter 07](../chapters/chapter-07.md) · 영역 2

## Q14. A

요청형 관리 실행 환경이 요구에 맞는다. 불필요한 cluster·물리 서버는 부담을 늘리고 DB만으로 API가 실행되지 않는다.

복습: [Chapter 17](../chapters/chapter-17.md) · 영역 2

## Q15. D

Kubernetes 요구와 관리형 운영의 조합이다. Cloud Run·메시지·NAT는 같은 API를 제공하지 않는다.

복습: [Chapter 19](../chapters/chapter-19.md) · 영역 2

## Q16. D

중단 감수형 설계로 Spot을 쓴다. 절대 가동 보장·백업 삭제·권한 확대는 조건에 맞지 않는다.

복습: [Chapter 09](../chapters/chapter-09.md) · 영역 2

## Q17. A

CPU/메모리 비율은 머신 사양 선택 영역이다. DNS·객체 등급·포트는 메모리 사양이 아니다.

복습: [Chapter 09](../chapters/chapter-09.md) · 영역 2

## Q18. D

Archive의 최소 기간과 조회 비용을 요구에 비교한다. 캐시·Local SSD는 이 보관 요구를 대체하지 않는다.

복습: [Chapter 10](../chapters/chapter-10.md) · 영역 2

## Q19. D

서로 다른 존의 배치가 필요하다. 동일 존·이름 변경은 장애 분리가 아니다.

복습: [Chapter 09](../chapters/chapter-09.md) · 영역 2

## Q20. C

파일 프로토콜 요구에 맞는 관리형 파일 서비스가 필요하다. 분석·메시지·DNS는 NFS가 아니다.

복습: [Chapter 10](../chapters/chapter-10.md) · 영역 2

## Q21. A

엔진 호환과 관리형 운영 조건에 부합한다. 나머지는 PostgreSQL 실행 엔진이 아니다.

복습: [Chapter 11](../chapters/chapter-11.md) · 영역 2

## Q22. C

요구를 Spanner의 관계형 분산 특성과 비교한다. 캐시·DNS·객체 이름은 같은 보장을 제공하지 않는다.

복습: [Chapter 11](../chapters/chapter-11.md) · 영역 2

## Q23. C

분석 쿼리에 맞는 서비스다. 나머지는 분석 저장·실행 엔진이 아니다.

복습: [Chapter 12](../chapters/chapter-12.md) · 영역 2

## Q24. D

wide-column·키 중심 요구를 Bigtable과 비교한다. DNS·빌드·NAT는 DB 대체품이 아니다.

복습: [Chapter 12](../chapters/chapter-12.md) · 영역 2

## Q25. A

메시지 전달·완충 목적에 맞는다. HA·백업·이미지 공개는 생산자-소비자 큐가 아니다.

복습: [Chapter 18](../chapters/chapter-18.md) · 영역 2

## Q26. D

피어링은 네트워크 연결 수단이다. 이름·DNS만으로 경로가 생기지 않고 저장 등급은 무관하다.

복습: [Chapter 15](../chapters/chapter-15.md) · 영역 2

## Q27. B

L4 내부 부하 분산 조건에 맞는 LB를 고른다. NAT·DNS·Lifecycle은 backend 부하 분산을 대체하지 못한다.

복습: [Chapter 14](../chapters/chapter-14.md) · 영역 2

## Q28. A

IPsec 터널 요구는 VPN이다. 저장·로그인·경보는 네트워크 터널이 아니다.

복습: [Chapter 15](../chapters/chapter-15.md) · 영역 2

## Q29. D, E

변경 계획 검토와 민감 파일 보호가 필요하다. state 공개·AI 무검토는 잘못이고 plan은 적용 자체가 아니다.

복습: [Chapter 20](../chapters/chapter-20.md) · 영역 2

## Q30. B

digest가 특정 이미지 내용을 고정한다. 변경 가능한 태그·이름·로그 삭제는 버전 고정이 아니다.

복습: [Chapter 16](../chapters/chapter-16.md) · 영역 2

## Q31. D

VM 상태와 앱 건강은 다르다. 결제·조직 변경은 원인 증거가 아니다.

복습: [Chapter 07](../chapters/chapter-07.md) · 영역 3

## Q32. C

이미지는 OS·소프트웨어 출발점을 재사용한다. 무작위 이미지·DNS·공개화는 일관된 배포가 아니다.

복습: [Chapter 08](../chapters/chapter-08.md) · 영역 3

## Q33. A

스냅샷 스케줄은 반복 디스크 백업 기능이다. 나머지는 복구 지점을 남기지 않는다.

복습: [Chapter 08](../chapters/chapter-08.md) · 영역 3

## Q34. D

가역적 트래픽 복구를 검토한다. 데이터 파괴·공개화·관측 제거는 복구가 아니다.

복습: [Chapter 17](../chapters/chapter-17.md) · 영역 3

## Q35. B

Pending은 배치 조건과 capacity를 확인한다. 아직 실행 안 된 앱 로그나 다른 제품 변경은 직접 원인이 아니다.

복습: [Chapter 19](../chapters/chapter-19.md) · 영역 3

## Q36. C

HPA는 수평 Pod 확장이다. VPA·백업·DNS는 Pod 수 자동 조절이 아니다.

복습: [Chapter 19](../chapters/chapter-19.md) · 영역 3

## Q37. C

객체 수명 주기 기능이 해당 작업을 수행한다. 이름·공개화·VM 확장은 무관하다.

복습: [Chapter 10](../chapters/chapter-10.md) · 영역 3

## Q38. C

논리적 오류에는 시점 복구가 필요하다. HA·CPU·인증 해제는 과거 상태 복원 수단이 아니다.

복습: [Chapter 11](../chapters/chapter-11.md) · 영역 3

## Q39. A

작업 상태와 오류 근거가 필요하다. 나머지는 job 실패의 직접 진단 정보가 아니다.

복습: [Chapter 12](../chapters/chapter-12.md) · 영역 3

## Q40. D

private DNS 범위 요구다. 공개 zone·이름 숨기기·공개 권한은 내부 해석 범위를 보장하지 않는다.

복습: [Chapter 14](../chapters/chapter-14.md) · 영역 3

## Q41. B

API 접근 특성에 맞는 경로를 검토한다. ingress·DNS 삭제·디스크는 발신 API 연결을 제공하지 않는다.

복습: [Chapter 14](../chapters/chapter-14.md) · 영역 3

## Q42. B

범위와 필터를 좁혀 사건을 조회한다. 나머지는 오류 로그 검색이 아니다.

복습: [Chapter 21](../chapters/chapter-21.md) · 영역 3

## Q43. A

라우팅과 쓰기 권한을 확인한다. 재시작·이메일·포맷은 근거 없는 조치다.

복습: [Chapter 21](../chapters/chapter-21.md) · 영역 3

## Q44. B, C

수집과 알림 전달을 각각 구성한다. 기본 metric·경보 자동 확장·cardinality 무영향 가정은 틀리다.

복습: [Chapter 22](../chapters/chapter-22.md) · 영역 3

## Q45. B

조건 평가와 알림 전달을 분리해 본다. 지표 삭제·복원·NAT 삭제는 채널 상태를 검증하지 않는다.

복습: [Chapter 22](../chapters/chapter-22.md) · 영역 3

## Q46. C

요청 경로별 지연은 tracing이 유용하다. 평균 하나·DNS·저장 등급은 경로별 증거가 아니다.

복습: [Chapter 23](../chapters/chapter-23.md) · 영역 3

## Q47. B

제공자 사건과 앱 영향을 대조한다. 나머지는 진단 증거를 없애거나 위험을 키운다.

복습: [Chapter 23](../chapters/chapter-23.md) · 영역 3

## Q48. C

세션·자원 수명과 브라우저 종료는 다르다. 출력·권한 변경은 실행 자원을 정리하지 않는다.

복습: [Chapter 23](../chapters/chapter-23.md) · 영역 3

## Q49. A

읽기 작업과 bucket 범위를 좁힌다. 나머지는 필요 이상 권한 또는 공개화다.

복습: [Chapter 24](../chapters/chapter-24.md) · 영역 4

## Q50. C

상위 부여는 상속된다. VM·DNS 변경은 IAM 상속을 제거하지 않는다.

복습: [Chapter 03](../chapters/chapter-03.md) · 영역 4

## Q51. A

서비스 계정 actAs는 별도 권한이다. 객체 읽기·DNS·결제 조회는 이를 제공하지 않는다.

복습: [Chapter 25](../chapters/chapter-25.md) · 영역 4

## Q52. C

단기 토큰 생성의 권한 요구를 확인한다. 나머지는 identity 토큰 발급 권한이 아니다.

복습: [Chapter 25](../chapters/chapter-25.md) · 영역 4

## Q53. B, E

federation과 좁은 신뢰 범위를 함께 설계한다. 키 공개·전면 신뢰·자원 공개는 요구에 맞지 않는다.

복습: [Chapter 25](../chapters/chapter-25.md) · 영역 4

## Q54. C

workload별 권한 분리가 맞다. 공유 키와 이미지 태그는 identity 분리가 아니다.

복습: [Chapter 25](../chapters/chapter-25.md) · 영역 4

## Q55. B

federation의 신뢰 범위를 필요한 workload로 좁힌다. 공급자 전체 신뢰·Owner·NAT는 요구에 맞지 않는다.

복습: [Chapter 25](../chapters/chapter-25.md) · 영역 4

## Q56. B

노출 비밀 회전과 저장 방식 수정이 필요하다. 태그 변경은 내용 제거가 아니며 공개화는 악화다.

복습: [Chapter 26](../chapters/chapter-26.md) · 영역 4

## Q57. D

암호화 의존성 원인을 조사한다. 공개화·CPU·TTL은 키 상태를 회복하지 않는다.

복습: [Chapter 26](../chapters/chapter-26.md) · 영역 4

## Q58. D

기존 역할과 실제 작업 요구를 비교한다. 나머지는 역할 설계 기준이 아니다.

복습: [Chapter 24](../chapters/chapter-24.md) · 영역 4

## Q59. B

계정은 principal이면서 정책을 가진 resource다. 무조건 동일시·삭제·공개화는 권한 모델을 설명하지 못한다.

복습: [Chapter 25](../chapters/chapter-25.md) · 영역 4

## Q60. B

정확한 거부 원인을 확인한다. 광범위 확대·관측 제거·공개화는 최소 권한 문제 해결 방식이 아니다.

복습: [Chapter 24](../chapters/chapter-24.md) · 영역 4
