# Blackjack Uygulaması

Bu proje, Go dilinde bir Blackjack oyunu geliştirmeyi ve onu bir web sunucusu olarak dağıtmayı amaçlamaktadır. Proje, kullanıcı yönetimi, masa ekleme, veritabanı bağlantıları, frontend geliştirme, güvenlik önlemleri ve yayına alma süreçlerini içerir.

## Proje Adımları

1. **Blackjack Oyunu Yazılacak**  
   Blackjack oyununun temel mantığı Go dilinde yazılacak. Oyuncu, krupiyeye karşı oynayacak ve klasik blackjack kuralları uygulanacak.  
   - [ ] Tamamlandı

2. **Auth Yazılacak**  
   Kullanıcı kimlik doğrulama sistemi eklenecek. Bu sistem, kullanıcıların giriş yapmasına, çıkış yapmasına ve kayıt olmasına izin verecek.  
   - [ ] Tamamlanmadı

3. **Kullanıcı ve Masa Eklenecek**  
   Birden fazla kullanıcı ve masa ekleme fonksiyonu geliştirilecek. Kullanıcılar farklı masalarda oturabilecek ve aynı anda birden fazla oyun oynayabilecek.  
   - [ ] Tamamlanmadı

4. **DB Bağlantıları Yapılacak**  
   Proje için gerekli olan veritabanı bağlantıları yapılacak. Bu, kullanıcı bilgilerini, oyun durumlarını ve masa bilgilerini saklamak için kullanılacak.  
   - [ ] Tamamlanmadı

5. **Frontend Eklenecek**  
   htmx ile basit bir frontend arayüzü geliştirilecek. Bu arayüz, oyuncunun oyun oynarken yapacağı işlemleri (hit, stand gibi) görsel olarak gösterecek.  
   - [ ] Tamamlanmadı

6. **Image ve Kod Güvenliği Sağlanacak**  
   Docker image'leri güvenlik taramasından geçirilecek ve kod güvenlik önlemleri alınacak. OWASP en iyi uygulamaları göz önünde bulundurulacak.  
   - [ ] Tamamlanmadı

7. **Yayına Alma**  
   Uygulama, Kubernetes üzerinde bir Load Balancer aracılığıyla yayına alınacak. Uygulamanın dış erişime açık olması sağlanacak.  
   - [ ] Tamamlanmadı


## Local Deneme

```bash
git clone <repo-url>
cd blackjack-app
go run main.go
localhost:8081
