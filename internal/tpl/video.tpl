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
