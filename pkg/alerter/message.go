package alerter

type Message struct {
	Embed []*MessageEmbed `json:"embeds,omitempty"`
}

// An MessageEmbed stores data for message embeds.
type MessageEmbed struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Color       int                `json:"color,omitempty"`
	Image       *MessageEmbedImage `json:"image,omitempty"`
}

type MessageEmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}
