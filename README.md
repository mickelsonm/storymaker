# storymaker
Scope:
--
For this project you will build a RESTful web application, which creates content-rich stories. A story is essentially a 
sequence of story blocks laid out on a timeline, which can contain imagery, text, video, or whatever can be displayed on
a webpage. Think of it as a web-based collage.

A story block is just an HTML5 structure, which can be placed anywhere on a timeline.

A collage can contain several story blocks or another sub timeline.


Example:
--
Suppose the story is about your life. You have a collage of story blocks about things that happened to you very early in life.

Example:
```xml
	<!-- timelines contain stories -->
	<timeline>
		<!-- stories contain collages -->
		<story>
			<collage>
				<story>
					<!-- stands alone -->
					<h1>Accepted the Nobel Peace Prize</h1>
				</story>
			</collage>
			<collage>
				<!-- a collage can have a sub timeline/sub-stories -->
				<timeline>
					<story>
						<collage>
							<h1>Graduated College</h1>
							<img>
							<img>
							<img>
						</collage>
					</story>
					<story>
						<collage>
							<h1>Bought our first house</h1>
							<img>
							<img>
							<img>
						</collage>
					</story>
				</timeline>
			</collage>
		</story>
	</timeline>
```

>> Note: This structure is not set in stone.

Purpose:
--
Who doesn't like to share or read a good story? A story-MAKER can share their stories.

Requirements:
--
You are building a web application that supports:
- Login/User management (user roles: Admin/Creator and user that logins to view or share)
- Uploading images and text for story blocks
- Building story blocks from those parts and assigning them to a collage (a collage could be a div or a separate html page)
- Assigning collages to a timeline (ability to order and give time duration)
- Edit/Save capability of story blocks, collage, or the entire story.
- Ability to archive/unarchive all of the pieces
- Ability to delete/undelete all of the pieces
- All pieces can be flagged as private, public, or shareable to specific users
- Appropriate views for CRUD of all the pieces
- Ability to swap content in the builder (images, text, etc)
- A live preview would be nice (Allows storyteller to see in real-time what they are building/what others would see)
- Ability to download stories (export as XML or JSON more than likely)
- Animations/carousel capability would be cool, but not necessarily required

Technologies:
--
- front-end (it would be fun to use): HTML5, CSS3, Javascript/AngularJS
- back-end (it would be fun to use): Golang, NodeJS, etc
- exports to: XML or JSON
