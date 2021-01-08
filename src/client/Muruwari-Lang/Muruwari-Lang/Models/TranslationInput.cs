using System;
using Newtonsoft.Json;

namespace Muruwari_Lang.Models
{
    public class TranslationInput
    {
        [JsonProperty("word")]
        public string WORD { get; set; }
    }
}
