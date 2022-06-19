<h1 align="center">Sei Chain</h1>

Ruesandora tarafından Türkçeleştirilmiştir.

![image](https://user-images.githubusercontent.com/101149671/172205075-389cfada-4e27-4c83-b2c8-0311b79868fa.png)

Sei Ağı, sipariş defterine özgü ilk L1 blok zinciridir. Zincir, her şeyden önce güvenilirliği, güvenliği ve yüksek verimi vurgulayarak, üstüne inşa edilmiş ultra yüksek performanslı DeFı ürünlerinin tamamen yeni bir kademesini mümkün kılıyor. Sei'nin zincir içi CLOB ve eşleştirme motoru, tüccarlar ve uygulamalar için derin likidite ve fiyat-zaman öncelikli eşleştirme sağlar. Seı üzerine kurulu uygulamalar, yerleşik sipariş defteri altyapısından, derin likiditeden ve tamamen merkezi olmayan bir eşleştirme hizmetinden yararlanır. Kullanıcılar, MEV koruması ile birlikte işlemlerinin fiyatını, boyutunu ve yönünü seçme yeteneği ile bu değişim modelinden yararlanır.

[Central limit order book (Clob) Nedir?](https://twitter.com/SeiTurkiye/status/1535687081221050368?s=20&t=dfa-2AbWgEdezAeJGWmL-Q)

# Sei chain

Sei Chain, Cosmos SDK ve Tender mint kullanılarak oluşturulmuş bir blockchain'dir. Cosmos SDK ve Tendermint çekirdeği kullanılarak oluşturulmuştur ve yerleşik bir merkezi limit sipariş defteri (CLOB) modülüne sahiptir. Sei'ye dayanan merkezi olmayan uygulamalar kulübün üzerine inşa edilebilir ve diğer Cosmos tabanlı blok zincirler, Sei'nin CLOB'UNU paylaşılan bir likidite merkezi olarak kullanabilir ve herhangi bir varlık için pazarlar oluşturabilir.

[Cosmos SDK Nedir?](https://github.com/ruesandora/Cosmos-SDK-TR)

Geliştiriciler ve kullanıcılar düşünülerek tasarlanan Sei, yeni nesil DeFı için altyapı ve paylaşılan likidite merkezi olarak hizmet vermektedir. Uygulamalar, Sei sipariş defteri altyapısında işlem yapmak ve diğer uygulamalardan havuzlanmış likiditeye erişmek için kolayca takılabilir ve oynatılabilir. Geliştirici deneyimine öncelik vermek için Sei Network, CosmWasm akıllı sözleşmelerini desteklemek üzere wsm modülünü entegre etti.

## Get started
**How to validate on the Sei Testnet**
*This is the Sei Testnet-1 (sei-testnet-1)*

Yukarıda ki kısmı çevirmek istemedim.

> Oluşum [Published](https://github.com/sei-protocol/testnet/blob/main/sei-testnet-1/genesis.json)

> Peers [Published](https://github.com/sei-protocol/testnet/blob/main/sei-testnet-1/addrbook.json)

## Donanım Gereksinimleri:
**Minimum**
* 8 GB RAM
* 100 GB NVME SSD
* 3.2 GHz x4 CPU

**Önerilen**
* 16 GB RAM
* 500 GB NVME SSD
* 4.2 GHz x6 CPU 

## İşletim sistemi:

> Linux (x86_64) veya Linux (amd64) Önerilen Arch Linux

**Sürüm**
> Ön koşul: go1.18+ gerekli.
* Arch Linux: `pacman -S go`
* Ubuntu: `sudo snap install go --classic`

> Ön koşul: git. 
* Arch Linux: `pacman -S git`
* Ubuntu: `sudo apt-get install git`

> İsteğe bağlı gereksinim: GNU make. 
* Arch Linux: `pacman -S make`
* Ubuntu: `sudo apt-get install make`

## Seid Kurulum Adımları

**Git deposunu klonla**

```bash
git clone https://github.com/sei-protocol/sei-chain
cd sei-chain
git checkout origin/1.0.1beta-upgrade
make install
mv $HOME/go/bin/seid /usr/bin/
```
**Anahtar oluştur**

* `seid keys add [key_name]`

* `seid keys add [key_name] --recover` to regenerate keys with your mnemonic

* `seid keys add [key_name] --ledger` to generate keys with ledger device

## Doğrulayıcı kurulum talimatları:

* seid ikili dosyasını yükleyi

* Düğümü başlat: `seid init <moniker> --chain-id sei-testnet-1`

* Genesis dosyasını indirin: `https://github.com/sei-protocol/testnet/raw/main/sei-testnet-1/genesis.json -P $HOME/.sei/config/`
 
* Asgari gaz fiyatlarını düzenleyin in ${HOME}/.sei/config/app.toml: `sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.01usei"/g' $HOME/.sei/config/app.toml`

* Düğümü arka planda çalıştırmak için bir systemd hizmeti oluşturarak seid'i başlatın
`nano /etc/systemd/system/seid.service`
> Aşağıdaki metni kopyalayıp servis dosyanıza yapıştırın. Uygun gördüğünüz gibi düzenlediğinizden emin olun.

```bash
[Unit]
Description=Sei-Network Node
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/root/
ExecStart=/root/go/bin/seid start
Restart=on-failure
StartLimitInterval=0
RestartSec=3
LimitNOFILE=65535
LimitMEMLOCK=209715200

[Install]
WantedBy=multi-user.target
```
## düğümü başlat:
* Servis dosyalarını yeniden yükleyin: `sudo systemctl daemon-reload` 
* symlink'i oluştur: `sudo systemctl enable seid.service` 
* sudo düğümünü başlat: `systemctl start seid && journalctl -u seid -f`

### Doğrulayıcı İşlemi Oluştur
```bash
seid tx staking create-validator \
--from {{KEY_NAME}} \
--chain-id  \
--moniker="<VALIDATOR_NAME>" \
--commission-max-change-rate=0.01 \
--commission-max-rate=1.0 \
--commission-rate=0.05 \
--details="<description>" \
--security-contact="<contact_information>" \
--website="<your_website>" \
--pubkey $(seid tendermint show-validator) \
--min-self-delegation="1" \
--amount <token delegation>usei \
--node localhost:26657
```
