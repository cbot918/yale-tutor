## Nginx Proxy with HTTPS Domain

適用: 對相關名詞要有點熟悉感，不然會不少查資料試誤

## 用到的技術: 
- [godaddy](https://tw.godaddy.com/) ( 買 domain，需費用，可以找一下有很便宜的，第一年一兩百那種 )
- [linode](https://www.linode.com/) ubuntu 20.04 ( 開機器，新手推薦開第二便宜那一台，一個月好像也是3百有找 )
- [golang](https://go.dev/doc/install) ( 寫 web server )
- html

## 組件:
- godaddy 買的 domain
- linode vm (後面簡稱機器)
- nginx proxy (機器內安裝)
- go server 兩台 (附程式碼)

## 流程

1. 開一台 linode 機器，測試一下 ssh 確保可登入，指令為 `ssh user@host`，其中 user 是使用者名稱，host 是 linode機器 的ip，ip 在 linode 網頁 dashboard可以看 ( 自己是用 vscode remote-ssh 去連機器, 好用大推)

2. linode 準備一下機器，ssh 進入 linode機器裡面之後

    2.1 安裝 nginx `sudo apt-get update && sudo apt-get install nginx`

    2.2 確認 nginx 狀況正常 `sudo systemctl status nginx`，看到 Active: active (running) 為正常

    2.3 cd 到 nginx 的設定資料夾內 `cd /etc/nginx` (我自己是會用 vscoded在 /etc/nginx 開一個視窗出來，方便很多)，介紹一下 nginx/sites-available 資料夾內是我們有的網站，nginx/sites-enabled 是打算開放出去的網站，目前應該會看到 default，這個是 nginx 預設頁面，此處我們可以先 `curl localhost` 會有看起來正常的網頁訊息，也可以用瀏覽器，輸入機器ip 去看。

    2.4 先做兩台 server 分別在 localhost:3000，localhost:3001，程式碼在servers.go，可以直接執行 `go run .`，運行兩台 server

3. godaddy 買一下 domain，怎麼挑 domain 可以去網路上找一下別人的心得，建議找個便宜又短的來練習，第一年都滿便宜的，以下假設我們買的網域為 `mytest777.com`

    3.1 買好之後，把 nameserver 從 godaddy 遷移到 linode 來管理 (實務上非必要，但這樣比較方便，後面的步驟也是基於這一步，所以以這份文件的角度來說是必要步驟)

    3.2 godaddy 找到自己的網域產品，進入[管理頁面]，進入[網域設定]，點[DNS]標籤，點[名稱伺服器]標籤 (這邊沒法很準確形容，但照著方向點會找到，按鈕細節也可以參考網路教學)

    3.3 在名稱伺服器，我們要從 godaddy 提供的，換成 [linode提供的](https://www.linode.com/docs/products/networking/dns-manager/guides/authoritative-name-servers/)
    ```
    ns1.linode.com
    ns2.linode.com
    ns3.linode.com
    ns4.linode.com
    ns5.linode.com
    ```
    3.4 點右邊中間[變更名稱伺服器], 開始做變更, domain 遷移會顯示最慢要 48hr，但自己兩次經驗 20分鐘內通常會完成，所以後面我們接著回 linode 做設定。

4. linode domain 做 A record 綁定，A record 簡單來說就是，把網址綁定我們的機器ip。

    4.1 點開 [linode domains](https://cloud.linode.com/domains)，點[Create Domain]，Domain輸入我們買的 `mytest777.com`，SOA Email Address輸入我們的email，Insert Default Records選 [Insert records from my one of linodes]，然後選目標機器。

    4.2 我們需要 `api.mytest777.com` 來測試看看，我們是否有做到 proxy 的功能，進入剛剛頁面我們已經新增的網址內，頁面找一個敘述 A/AAAA Record 右邊按鈕 [Add An A/AAAA Record]點下去，Hostname 輸入 api (其實就是 api.mytest777.com的效果)，IP Address 輸入我們機器的 IP，按 [Save]就ok了

    4.3 [看domain綁定機器有沒有成功的網站](https://www.whatsmydns.net/#A/blog.yale.codes)，輸入 `mytest777.com` 及 `api.mytest777.com`，去看有沒有通，經驗上是5分鐘以內會通。

5. linode nginx 寫 proxy 設定檔

    5.1 回到機器內，/etc/nginx

    5.2 在 /etc/nginx/sites-available，加入 `mytest777.com` 及 `api.mytest777.com` 這兩個檔案，檔名只是好管理，重點在檔案內容，兩行註解想懂，在想一下 proxy server 是監聽80 port，觀念就會比較清楚



    5.3 因為 sites-available 是設定倉庫，sites-enabled 是現在要開出去的，所以用個 linux 提供的軟連結功能，在sites-enabled裡面的設定檔會被 nginx 執行 (軟連結就是，我們只會修改 sites-available 裡面的內容，另一邊就會跟著更新)
    ```
        sudo ln -s /etc/nginx/sites-available/mytest777.com /etc/nginx/sites-enabled/

        sudo ln -s /etc/nginx/sites-available/api.mytest777.com /etc/nginx/sites-enabled/
    ```
    然後我會把 sites-enabled 裡面的 default 刪掉，因為自己實作的時候，這個檔案有影響到結果，刪掉後結果才會是預期

    5.4 先機器內做個測試
    ```bash
      curl mytest777.com #回應 server in port 3000

      curl api.mytest777.com ##回應 server in port 3001
    ```

    5.5 正式網頁測試
    ```bash
    # 瀏覽器分別輸入
    http://mytest777.com
    http://api.mytest777.com
    # 預期看到對應回應文字
    ```

6. 幫網址加上 TSL 功能

    6.1 觀念：TSL需要 webserver 本身有證書，證書又需要申請，在此我們使用免費證書商 `Let's encrypt`，每三個月需要重新認證，所以使用自動化申請續約套件 `Certbot`

    6.2 安裝 Certbot
    ```bash
    sudo apt install certbot python3-certbot-nginx -y
    ```
    
    6.3 申請證書，先目標 `mytest777.com` 這個網址，其他細節就看自己決定啦
    ```bash
    sudo certbot --nginx
    ```

    6.4 驗證有沒有錯誤
    ```bash
      sudo systemctl status certbot.timer
      sudo certbot renew --dry-run
    ```

    6.5 測試 https    
    ```bash
    # 瀏覽器分別輸入
    https://mytest777.com
    # 終於有鎖頭受保護啦
    ```

## 參考:
[Ubuntu：安裝 Nginx 做反向代理（Reverse Proxy）設定教學](https://mnya.tw/cc/word/1921.html)

[GoDaddy 域名 + Linode 主機商 + Https 憑證 + 自動更新](https://tasb00429.medium.com/godaddy-%E5%9F%9F%E5%90%8D-linode-%E4%B8%BB%E6%A9%9F%E5%95%86-https-%E6%86%91%E8%AD%89-3c2273189725)

[Connect a Domain to a Linode Server](https://www.youtube.com/watch?v=mKfx4ryuMtY)

[how to get free ssl](https://www.tutorialsteacher.com/https/get-free-ssl-certificate)


## 延伸閱讀:
[how-to-secure-nginx-with-let-s-encrypt-on-ubuntu](https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-20-04)


## 後記：

自學 infra 還是滿複雜且累的，這一篇算是比較統整性的視角去學習。

文章實作遇到問題，歡迎來 [DC](https://discord.gg/GwJcrhPT7h)找我唷

文末想要推廣一下，自動化佈署平台 [Zeabur](https://zeabur.com)，非常好用！之後有機會再來介紹