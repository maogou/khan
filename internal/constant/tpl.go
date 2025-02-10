package constant

const SnsText = `
<TimelineObject>
    <id></id>
    <username></username>
    <createTime></createTime>
    <contentDesc><![CDATA[{{.Content}}]]></contentDesc>
    <contentDescShowType>0</contentDescShowType>
    <contentDescScene>3</contentDescScene>
    <private>0</private>
    <sightFolded>0</sightFolded>
    <showFlag>0</showFlag>
    <appInfo>
        <id></id>
        <version></version>
        <appName></appName>
        <installUrl></installUrl>
        <fromUrl></fromUrl>
        <isForceUpdate>0</isForceUpdate>
        <isHidden>0</isHidden>
    </appInfo>
    <sourceUserName></sourceUserName>
    <sourceNickName></sourceNickName>
    <statisticsData></statisticsData>
    <statExtStr></statExtStr>
    <ContentObject>
        <contentStyle>2</contentStyle>
        <title></title>
        <description></description>
        <mediaList></mediaList>
        <contentUrl></contentUrl>
    </ContentObject>
    <actionInfo>
        <appMsg>
            <messageAction></messageAction>
        </appMsg>
    </actionInfo>
    <location poiClassifyId="" poiName="" poiAddress="" poiClassifyType="0" city=""></location>
    <publicUserName></publicUserName>
    <streamvideo>
        <streamvideourl></streamvideourl>
        <streamvideothumburl></streamvideothumburl>
        <streamvideoweburl></streamvideoweburl>
    </streamvideo>
</TimelineObject>
`

const SnsImage = `
<TimelineObject>
    <id>0</id>
    <username></username>
    <createTime>1603345160</createTime>
    <contentDesc><![CDATA[${content}]]></contentDesc>
    <contentDescShowType>0</contentDescShowType>
    <contentDescScene>3</contentDescScene>
    <private>0</private>
    <sightFolded>0</sightFolded>
    <showFlag>0</showFlag>
    <contentattr><![CDATA[0]]></contentattr>
    <sourceUserName></sourceUserName>
    <sourceNickName></sourceNickName>
    <statisticsData></statisticsData>
    <weappInfo>
        <appUserName></appUserName>
        <pagePath></pagePath>
        <version><![CDATA[0]]></version>
        <isHidden>0</isHidden>
        <debugMode><![CDATA[0]]></debugMode>
        <shareActionId></shareActionId>
        <isGame><![CDATA[0]]></isGame>
        <messageExtraData></messageExtraData>
        <subType><![CDATA[0]]></subType>
        <preloadResources></preloadResources>
    </weappInfo>
    <ContentObject>
        <contentStyle><![CDATA[1]]></contentStyle>
        <contentSubStyle><![CDATA[0]]></contentSubStyle>
        <title></title>
        <description></description>
        <contentUrl></contentUrl>
        <mediaList><#list imgInfos as img>
                <media>
                <id>0</id>
                <type>2</type>
                <title></title>
                <private>0</private>
                <userData></userData>
                <subType>0</subType>
                <videoSize width="500" height="889"/>
                <url type="1" <#if img.fileMd5?? > md5="${img.fileMd5}"</#if>><![CDATA[${img.fileUrl}]]></url>
                <thumb type="1"><![CDATA[${img.thumbUrl}]]></thumb></media></#list></mediaList>
    </ContentObject>
    <actionInfo>
        <appMsg>
            <mediaTagName></mediaTagName>
            <messageExt></messageExt>
            <messageAction></messageAction>
        </appMsg>
    </actionInfo>
    <appInfo>
        <id></id>
    </appInfo>
    <publicUserName></publicUserName>
    <streamvideo>
        <streamvideourl></streamvideourl>
        <streamvideothumburl></streamvideothumburl>
        <streamvideoweburl></streamvideoweburl>
    </streamvideo>
</TimelineObject>
`

const SnsUrl = `
<TimelineObject>
    <id>0</id>
    <username><![CDATA[]]></username>
    <createTime>1604548459</createTime>
    <contentDesc><![CDATA[${contentDesc}]]></contentDesc>
    <contentDescShowType>0</contentDescShowType>
    <contentDescScene>4</contentDescScene>
    <private>0</private>
    <sightFolded>0</sightFolded>
    <showFlag>0</showFlag>
    <appInfo>
        <id></id>
        <version></version>
        <appName></appName>
        <installUrl></installUrl>
        <fromUrl></fromUrl>
        <isForceUpdate>0</isForceUpdate>
    </appInfo>
    <sourceUserName></sourceUserName>
    <sourceNickName></sourceNickName>
    <statisticsData></statisticsData>
    <statExtStr></statExtStr>
    <ContentObject>
        <contentStyle>3</contentStyle>
        <title><![CDATA[${title}]]></title>
        <description><![CDATA[${description}]]></description>
        <contentUrl><![CDATA[${contentUrl}]]></contentUrl>
        <mediaList>
            <media>
                <id>0</id>
                <type>2</type>
                <title></title>
                <description></description>
                <private>0</private>
                <userData></userData>
                <subType>0</subType>
                <videoSize width=\"0\" height=\"0\"></videoSize>
                <url type=\"0\" md5=\"\" videomd5=\"\"><![CDATA[${thumbUrl}]]></url>
                <lowBandUrl type=\"0\"><![CDATA[${thumbUrl}]]></lowBandUrl>
                <thumb type=\"1\"><![CDATA[${thumbUrl}]]></thumb>
            </media>
        </mediaList>
        <mmreadershare>
            <itemshowtype>0</itemshowtype>
            <nativepage>0</nativepage>
            <digest>
                <![CDATA[]]>
            </digest>
            <cover>
                <![CDATA[]]>
            </cover>
            <pubtime>0</pubtime>
            <duration>0</duration>
            <width>0</width>
            <height>0</height>
            <vid></vid>
            <funcflag>0</funcflag>
        </mmreadershare>
    </ContentObject>
    <actionInfo>
        <appMsg>
            <messageAction></messageAction>
        </appMsg>
    </actionInfo>
    <location poiClassifyId=\"\" poiName=\"\" poiAddress=\"\" poiClassifyType=\"0\" city=\"\"></location>
    <publicUserName></publicUserName>
    <streamvideo>
        <streamvideourl></streamvideourl>
        <streamvideothumburl></streamvideothumburl>
        <streamvideoweburl></streamvideoweburl>
    </streamvideo>
</TimelineObject>
`

const SnsVideo = `
<TimelineObject>
    <id>0</id>
    <username><![CDATA[]]></username>
    <createTime>1603346167</createTime>
    <contentDesc><![CDATA[${content}]]></contentDesc>
    <contentDescShowType>0</contentDescShowType>
    <contentDescScene>0</contentDescScene>
    <private>0</private>
    <sightFolded>0</sightFolded>
    <showFlag>0</showFlag>
    <appInfo>
        <id></id>
        <version></version>
        <appName></appName>
        <installUrl></installUrl>
        <fromUrl></fromUrl>
        <isForceUpdate>0</isForceUpdate>
    </appInfo>
    <sourceUserName></sourceUserName>
    <sourceNickName></sourceNickName>
    <statisticsData></statisticsData>
    <statExtStr></statExtStr>
    <ContentObject>
        <contentStyle>15</contentStyle>
        <title></title>
        <description>Sight</description>
        <mediaList>
            <media>
                <id>0</id>
                <type>6</type>
                <title></title>
                <description><![CDATA[${content}]]></description>
                <private>0</private>
                <userData></userData>
                <subType>0</subType>
                <videoSize width=\"720\" height=\"1280\"/>
                    <url type=\"1\">
                        <![CDATA[${videoInfo.fileUrl}]]>
                    </url>
                    <thumb type=\"1\"><![CDATA[${videoInfo.thumbUrl}]]></thumb>
                    <size width=\"720.000000\" height=\"1280.000000\" totalSize=\"1649264\"/>
                        <videoDuration>
                            <#if videoInfo.videoDuration?? >
                                <![CDATA[${videoInfo.videoDuration}]]>
                            </#if>
                        </videoDuration>
            </media>
        </mediaList>
    </ContentObject>
</TimelineObject>
`
