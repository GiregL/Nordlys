namespace Nordlys.Data;

/// <summary>
/// Currency conversion (based on Euro)
/// </summary>
public class CurrencyConversion
{
    public int Id { get; set; }
    public DateTime ConversionDate { get; set; }
    
    public int TargetId { get; set; }
    public Currency Target { get; set; }
}