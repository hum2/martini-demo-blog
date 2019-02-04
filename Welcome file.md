＃StackEditへようこそ！

私は** StackEdit **の最初のMarkdownファイルです。あなたがStackEditについて学びたいのなら、あなたは私を読んでということです。またあなたはMarkdownで遊びたいのですが、私は編集私の仕事です終わったら、ナビゲーションバーの左隅にある**ファイルエクスプロラ**を開いて新しいファイルを作成できます。


＃ファイル

StackEditはあなたのファイルをあなたのブラウザに保存します。まとめ、あなたのすべてのファイルは自動的にローカルに保存され、**オフラインでアクセス可能になります。**

## ファイルとフォルダを作成する

ファイルエクスプローラには、ナビゲーションバーの左隅にあるボタンを使用してアクセスできます。ファイルエクスプローラの**新しいファイル**ボタンをクリックして新しいファイルを作成できます。**新しいフォルダ**ボタンをクリックフォルダを作成することもできます。

## 他のファイルに切り替え

すべてのファイルはファイルエクスプロラに表示されています。リスト内のファイルをクリックすると、別のものに切り替えすることができます。

## ファイル名を変更する

ナビゲーションバーのファイル名をクリックするか、ファイルエクスプローラの**名前の変更**ボタンをクリックして、現在のファイルの名前を変更できます。

## ファイルを削除する

ファイルは** Trash **フォルダに移動さ​​れ、7日間操作がないと自動的に削除されます。ファイルエクスプロラの**削除**ボタンをクリックしてください。

## ファイルをエキスパートする

現在のファイルをエクスポートするには、メニューの**ディスクにエクスポート**をクリックします。ファイルをプレーンな値下げとして、ハンドルバーテンプレートを使用したHTMLとして、またはPDFとしてエクスポートすることを選択できます。


＃同期

Synchronization is one of the biggest features of StackEdit. It enables you to synchronize any file in your workspace with other files stored in your **Google Drive**, your **Dropbox** and your **GitHub** accounts. This allows you to keep writing on other devices, collaborate with people you share the file with, integrate easily into your workflow... The synchronization mechanism takes place every minute in the background, downloading, merging, and uploading file modifications.

There are two types of synchronization and they can complement each other:

- The workspace synchronization will sync all your files, folders and settings automatically. This will allow you to fetch your workspace on any other device.
	> To start syncing your workspace, just sign in with Google in the menu.

- The file synchronization will keep one file of the workspace synced with one or multiple files in **Google Drive**, **Dropbox** or **GitHub**.
	> Before starting to sync files, you must link an account in the **Synchronize** sub-menu.

## Open a file

You can open a file from **Google Drive**, **Dropbox** or **GitHub** by opening the **Synchronize** sub-menu and clicking **Open from**. Once opened in the workspace, any modification in the file will be automatically synced.

## Save a file

You can save any file of the workspace to **Google Drive**, **Dropbox** or **GitHub** by opening the **Synchronize** sub-menu and clicking **Save on**. Even if a file in the workspace is already synced, you can save it to another location. StackEdit can sync one file with multiple locations and accounts.

## Synchronize a file

Once your file is linked to a synchronized location, StackEdit will periodically synchronize it by downloading/uploading any modification. A merge will be performed if necessary and conflicts will be resolved.

If you just have modified your file and you want to force syncing, click the **Synchronize now** button in the navigation bar.

> **Note:** The **Synchronize now** button is disabled if you have no file to synchronize.

## Manage file synchronization

Since one file can be synced with multiple locations, you can list and manage synchronized locations by clicking **File synchronization** in the **Synchronize** sub-menu. This allows you to list and remove synchronized locations that are linked to your file.


# Publication

Publishing in StackEdit makes it simple for you to publish online your files. Once you're happy with a file, you can publish it to different hosting platforms like **Blogger**, **Dropbox**, **Gist**, **GitHub**, **Google Drive**, **WordPress** and **Zendesk**. With [Handlebars templates](http://handlebarsjs.com/), you have full control over what you export.

> Before starting to publish, you must link an account in the **Publish** sub-menu.

## Publish a File

You can publish your file by opening the **Publish** sub-menu and by clicking **Publish to**. For some locations, you can choose between the following formats:

- Markdown: publish the Markdown text on a website that can interpret it (**GitHub** for instance),
- HTML: publish the file converted to HTML via a Handlebars template (on a blog for example).

## Update a publication

After publishing, StackEdit keeps your file linked to that publication which makes it easy for you to re-publish it. Once you have modified your file and you want to update your publication, click on the **Publish now** button in the navigation bar.

> **Note:** The **Publish now** button is disabled if your file has not been published yet.

## Manage file publication

Since one file can be published to multiple locations, you can list and manage publish locations by clicking **File publication** in the **Publish** sub-menu. This allows you to list and remove publication locations that are linked to your file.


# Markdown extensions

StackEdit extends the standard Markdown syntax by adding extra **Markdown extensions**, providing you with some nice features.

> **ProTip:** You can disable any **Markdown extension** in the **File properties** dialog.


## SmartyPants

SmartyPants converts ASCII punctuation characters into "smart" typographic punctuation HTML entities. For example:

|                |ASCII                          |HTML                         |
|----------------|-------------------------------|-----------------------------|
|Single backticks|`'Isn't this fun?'`            |'Isn't this fun?'            |
|Quotes          |`"Isn't this fun?"`            |"Isn't this fun?"            |
|Dashes          |`-- is en-dash, --- is em-dash`|-- is en-dash, --- is em-dash|


## KaTeX

You can render LaTeX mathematical expressions using [KaTeX](https://khan.github.io/KaTeX/):

The *Gamma function* satisfying $\Gamma(n) = (n-1)!\quad\forall n\in\mathbb N$ is via the Euler integral

$$
\Gamma(z) = \int_0^\infty t^{z-1}e^{-t}dt\,.
$$

> You can find more information about **LaTeX** mathematical expressions [here](http://meta.math.stackexchange.com/questions/5020/mathjax-basic-tutorial-and-quick-reference).


## UML diagrams

You can render UML diagrams using [Mermaid](https://mermaidjs.github.io/). For example, this will produce a sequence diagram:

` `  `マーメイド
シーケンス図
アリス- >>ボブ：こんにちはボブ、お元気ですか？
ボブ- >>ジョン：ジョンはどうですか？
ボブ-アリス×：ありがとうございます
。ボブ-Xジョン：ありがとうございます
ジョンの右記：ボブは長い<br/>長い時間を考えているので、テキストは1行に収まりません。ボブ - >アリス：ジョンと確認しています…アリス - >ジョン：はい…ジョン、お元気ですか？` ` `





それはこれはフローチャートを作成します：

` `  `人魚
ブログラフLR 
A [四角] -リンクテキスト- > B（（丸））
A - > C（丸四角形）
B - > D {菱形} 
C - > D `  ` `

<!--stackedit_data:
eyJoaXN0b3J5IjpbLTU1MDgwOThdfQ==
-->