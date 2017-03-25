---
layout: post
title: 如何使用Hexo和Github搭建自己的Blog
date: 2015-06-22 09:58:17
tags:
- Hexo
- Jacman
- Github
categories: Hexo
description: The tutoria will show you how to set up a Hexo site and deply it to Github Page.
---
##主题介绍
大多数的Programer都会有写Blog的习惯。使用怎样的Blog平台呢？大概有这么几种：
1. sina，sohu等门户网站都提供来blog服务。
2. 在cnblogs，csdn或iteye等比较专业的IT社区上建立blog
3. 自己购买虚拟主机，数据库和域名。自己开发blog站点，并部署到自己的server上
4. Github提供了静态Blog的托管服务。

上面的4种方案，我来简单的做一下比较吧：
1. 不做评论，直接pass。上大学的时候我就在sina上写blog，现在的小学生都有自己的sina blog了。
2. 我在cnblogs上也有自己的blog，cnblog有这么些缺陷：很难定制化，url不友好。最关键是有广告。
3. 将站点托管在自己的server上挺不错的，但需要花费money。
4. 作为全球最大的programer(男性)交友平台，将blog部署到github上是一个不错的选择。

我的Blog就是host到github上的。我将blog托管到github上有如下原因：
1. github支持静态blog的托管，是完全免费的。
2. github是全球最大最稳定的Git托管服务提供商，你不用担心什么时候你的blog不可用了。
3. 无广告。
4. 你拥有blog站点的完全控制权。
5. Github支持为你的blog绑定域名。

<!-- more -->

Github只支持托管静态的站点。静态站点是相对于动态站点来说的，简单说就是没有数据库的站点，只有一些静态资源，比如html，css，javascript等。其实动态站点最终呈现的也是html。静态站点比较适合做blog。而且静态blog相对于动态blog有一个最大的好处：速度快。
创建一个静态blog有2种方法：
1.从零开始，完全自己去写。
2.使用一些静态blog生成程序。
现在有很多的静态blog framework，比如：Jekyll，[Hexo](http://http://hexo.io/)，Octopress等。使用这些静态blog框架可以很快的就创建一个blog站点。Github使用的静态blog engine是Jekyll，所以你在Github上创建一个blog有2种方法：将生成好的Blog站点部署到Github。将Jekyll的源文件托管到Github上，Github会自动替你解析成html。
我最早是使用Jekyll的，但有些地方不太方便，现在我使用的是Hexo。Hexo比Jekyll更简单，速度更快。

我使用的Hexo主题是[Jacman](https://github.com/wuchong/jacman)，一款扁平化，有响应式设计的主题。自己做了些修改。并添加了站内搜索功能和评论功能。自己还是比较满意的。下面就讲自己从零开始搭建Hexo站点的过程记录下来吧。

## Installation Instruction
### Install Hexo
Hexo is based on Node.js, before installing Hexo, Node.js and Git are required.
#### Install Node.js
**By Homebrew**: You can use Homebrew to install Node.js. Type the command `brew install node` will install Node.js on your machine. If you encounter the *No Permission* error, please use `sudo brew install node`.    
**From Installation Package**: Due to the newwork limit, it may be very slow to downlaod Node.js by Homebrew. As an alternative, you can download Node.js from official website and install it.
#### Install Git
Use `sudo brew install git` to install Git.
#### Install Hexo
After you install Node.js successfully，you can use the `npm` command to install Hexo。 The `npm` tool come with Node.js, out of box.
Run the following command to install Hexo.
`sudo npm install -g hexo-cli`    

Note: the Mac user please kinp in mind，some error may occurred while compiling. Please first download and install XCode. XCode is the size of 2.5G，you can get it from Apple  App Store。下载好后，打开XCode，接受license，XCode会下载一些组件，初始化。
请确保npm在PATH变量中。
`vi ~/.bash_profile`
`export PATH=/usr/local/bin/npm:PATH`
Type command 'hexo version' to check if the Hexo is installed successfully.

### Setup your site
Once Hexo is installed, run the following commands to initialize a Hexo site in a target <folder>
```bash
hexo init <folder>  # Hexo will create a site template in the target folder.
```
Once initialized, the following is what your site folder will look like:
```
.
├── _config.yml       # The site configuration file. It is in the yaml format.
├── package.json      # Specify the dependency node.js modules.
├── scaffolds
├── scripts
├── source
|   ├── _drafts       # Contains your post draft. the drafts will not display in the site.
|   └── _posts        # Contains your posts of your site.
└── themes            # Contains all the themes you installed.
```
For more details of the folder in the above diagram, please visit [the Hexo official document of this section](https://hexo.io/docs/setup.html).

Pay attention to the package.json file, it includes the dependencies of your site, so before starting your site, you need to install the dependencies node.js modules.

execute below command to install dependencies:
```bash
cd <folder> # cd to the target foldser.
npm install # Hexo will install the node.js modules specified in the package.json.
```
It may take a few minutes. As Hexo need to download the packages from remote server. How long it take depends on the situation of your network.

### Run your site at local server
Now, you can run your site locally by following steps:
1. Execute `hexo generate`, the `generate` command will generate your static site, and put all the files to a `public` folder in your site root folder.
2. Execute `hexo server`, the `server` command will start a web server, and run your site. The command will tell you the url of your site. The default port is 4000. For example, `http://localhost:4000`.
3. Visit the url in your web browser.

If you can visit it, congratulate to you, you have set up your site successfully. All the files of your static website are located in the `public` folder, so you can just copy all the files in `public` folder to Github Page. Don't care about how to set up the [Github Page](https://pages.github.com/), I will show you in a minute.

## Set up a Github Page
Github Page supply the service of hosting a satic website in it. You can take below steps to do it:
1. Create a Github account.
2. Create the Github Page repository.
3. Configure your SSH key to connect to Github.
4. Push your website to the master branch of your repository.

### Create a Github account
Register a Github account is totally free.
### Create the Github Page repository
Create a repository whose name is exactly `<your_github_user_name>.github.io`. For example, my github account is `leo-lei`, and my github repository of the site is `leo-lei.github.io`.
### Configure your SSH key
Github has several transfer protocal you can use to connect to your Github repository. For example, HTTPS(https://github.com/Leo-Lei/Leo-Lei.github.io.git), SSH(git@github.com:Leo-Lei/Leo-Lei.github.io.git) and SVN(https://github.com/Leo-Lei/Leo-Lei.github.io). I strongly recommended you to use the SSH protocal.
#### Generating SSH keys
Setp 1: Check for SSH keys
```bash
ls -la ~/.ssh
```
To check if there is a `id_rsa.pub` file.
Setp 2 : Generate a new SSH key
```bash
ssh-keygen
```
The teminator will ask you for the password, you can just press `Enter` to not use password.
The command will generate a SSH public key and private key. The private key is only for you, please don't tell it to others.
Setp 3: Add the SSH public key to your Github account
You need to add the SSH public key to Github by below steps:
1. Go to: Settings => SSH keys => Add SSH key
2. Copy and paste your public key to the key field.

If you want to use github on multiple computers, you need to generate a SSH key for each machine and add all of them to Github.

### Install a Hexo theme
Hexo generate a static website by combining the site content with the theme. There are some themes with Hexo distribution out of box, meanwhile Hexo support creating your own theme. My site use the Hexo theme [Jacman](https://github.com/wuchong/jacman). The author of the theme is a former student of the college.

Execute below command in the root folder of your site:
```bash
$ git clone https://github.com/wuchong/jacman.git themes/jacman
```
